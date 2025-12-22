package wifi

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mockErrInterfaces = errors.New("failed to retrieve interfaces")
)

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	macAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	macAddr2, _ := net.ParseMAC("66:77:88:99:aa:bb")

	expectedInterfaces := []*wifi.Interface{
		{HardwareAddr: macAddr1},
		{HardwareAddr: macAddr2},
	}

	mockHandle.On("Interfaces").Return(expectedInterfaces, nil)

	svc := New(mockHandle)

	addrs, err := svc.GetAddresses()

	require.NoError(t, err)
	assert.Len(t, addrs, 2)
	assert.Equal(t, []net.HardwareAddr{macAddr1, macAddr2}, addrs)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return(nil, mockErrInterfaces)

	svc := New(mockHandle)

	result, err := svc.GetAddresses()

	require.Error(t, err)
	assert.ErrorContains(t, err, "getting interfaces")
	assert.Nil(t, result)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return([]*wifi.Interface{}, nil)

	svc := New(mockHandle)

	addrs, err := svc.GetAddresses()

	require.NoError(t, err)
	assert.Empty(t, addrs)
	assert.NotNil(t, addrs)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	expectedInterfaces := []*wifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}

	mockHandle.On("Interfaces").Return(expectedInterfaces, nil)

	svc := New(mockHandle)

	names, err := svc.GetNames()

	require.NoError(t, err)
	assert.Equal(t, []string{"wlan0", "eth0"}, names)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewMockWiFiHandle(t)

	mockHandle.On("Interfaces").Return(nil, mockErrInterfaces)

	svc := New(mockHandle)

	names, err := svc.GetNames()

	require.Error(t, err)
	assert.ErrorContains(t, err, "getting interfaces")
	assert.Nil(t, names)
}
