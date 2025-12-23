package wifi_test

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	myWiFi "github.com/denisK-H/task-6/internal/wifi"
)

const (
	iface1Name = "wlan0"
	iface2Name = "wlan1"
)

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mac1 := net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	mac2 := net.HardwareAddr{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifi.Interface{
		{
			Name:         iface1Name,
			HardwareAddr: mac1,
		},
		{
			Name:         iface2Name,
			HardwareAddr: mac2,
		},
	}, nil)

	service := myWiFi.New(mockHandle)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	require.Len(t, addrs, 2)

	assert.Equal(t, mac1, addrs[0])
	assert.Equal(t, mac2, addrs[1])

	mockHandle.AssertExpectations(t)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return(nil, assert.AnError)

	service := myWiFi.New(mockHandle)

	addrs, err := service.GetAddresses()
	require.Nil(t, addrs)
	require.Error(t, err)
	require.ErrorContains(t, err, "getting interfaces")

	mockHandle.AssertExpectations(t)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return([]*wifi.Interface{
		{Name: iface1Name},
		{Name: iface2Name},
	}, nil)

	service := myWiFi.New(mockHandle)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{iface1Name, iface2Name}, names)

	mockHandle.AssertExpectations(t)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	mockHandle.On("Interfaces").Return(nil, assert.AnError)

	service := myWiFi.New(mockHandle)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.Error(t, err)
	require.ErrorContains(t, err, "getting interfaces")

	mockHandle.AssertExpectations(t)
}
