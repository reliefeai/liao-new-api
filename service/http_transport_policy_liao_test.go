package service

import (
	"testing"

	"github.com/QuantumNous/new-api/relaykit/dto"

	"github.com/stretchr/testify/assert"
)

// liao 旧开关 force_http1 必须继续生效，等价于 http_protocol=http1。
func TestNormalizeHTTPTransportPolicyLegacyForceHTTP1(t *testing.T) {
	p := NormalizeHTTPTransportPolicy(dto.ChannelSettings{ForceHTTP1: true})
	assert.Equal(t, dto.HTTPProtocolHTTP1, p.Protocol)
	assert.Equal(t, 1, p.Shards)

	p = NormalizeHTTPTransportPolicy(dto.ChannelSettings{ForceHTTP1: true, HTTPProtocol: "auto", HTTP2ConnectionShards: 4})
	assert.Equal(t, dto.HTTPProtocolHTTP1, p.Protocol)
	assert.Equal(t, 1, p.Shards)

	p = NormalizeHTTPTransportPolicy(dto.ChannelSettings{})
	assert.Equal(t, dto.HTTPProtocolAuto, p.Protocol)
}
