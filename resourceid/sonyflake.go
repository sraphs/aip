package resourceid

import (
	"errors"
	"hash/fnv"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

type sonyflakeGenerator struct {
	*sonyflake.Sonyflake
}

func (s *sonyflakeGenerator) Next() (string, error) {
	id, err := s.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}

var (
	sonyFlakeGenerator *Generator = nil
)

func SonyFlakeGenerator() Generator {
	if sonyFlakeGenerator == nil {
		sfg := Generator(&sonyflakeGenerator{
			sonyflake.NewSonyflake(sonyflake.Settings{
				MachineID: machineID,
				StartTime: time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC),
			}),
		})

		sonyFlakeGenerator = &sfg
	}

	return *sonyFlakeGenerator
}

func machineID() (uint16, error) {
	if mID, err := hostname(); err == nil {
		return mID, nil
	}

	if mID, err := lower16BitPrivateIP(); err == nil {
		return mID, nil
	}

	return 0, nil
}

// the following is a copy of sonyflake (https://github.com/sony/sonyflake/blob/master/sonyflake.go)
func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func hostname() (uint16, error) {
	host, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	h := fnv.New32()
	_, hashErr := h.Write([]byte(host))
	if hashErr != nil {
		return 0, hashErr
	}

	return uint16(h.Sum32()), nil
}
