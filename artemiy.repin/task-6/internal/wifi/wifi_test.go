package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/Nevermind0911/task-6/internal/wifi"
	mdlayherwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var errInterfaces = errors.New("failed to retrieve interfaces")

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	macAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	macAddr2, _ := net.ParseMAC("66:77:88:99:aa:bb")

	expectedInterfaces := []*mdlayherwifi.Interface{
		{HardwareAddr: macAddr1},
		{HardwareAddr: macAddr2},
	}

	mockHandle.On("Interfaces").Return(expectedInterfaces, nil)

	svc := wifi.New(mockHandle)

	addrs, err := svc.GetAddresses()

	require.NoError(t, err)
	require.Len(t, addrs, 2)
	require.Equal(t, []net.HardwareAddr{macAddr1, macAddr2}, addrs)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return(nil, errInterfaces)

	svc := wifi.New(mockHandle)

	result, err := svc.GetAddresses()

	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, result)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return([]*mdlayherwifi.Interface{}, nil)

	svc := wifi.New(mockHandle)

	addrs, err := svc.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, addrs)
	require.NotNil(t, addrs)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	expectedInterfaces := []*mdlayherwifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}

	mockHandle.On("Interfaces").Return(expectedInterfaces, nil)

	svc := wifi.New(mockHandle)

	names, err := svc.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth0"}, names)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return(nil, errInterfaces)

	svc := wifi.New(mockHandle)

	names, err := svc.GetNames()

	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, names)
}
