package fuzz

import (
	"testing"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
	"github.com/envoyproxy/gateway/internal/envoygateway/config"
	"github.com/envoyproxy/gateway/internal/logging"
	"github.com/stretchr/testify/require"
)

func FuzzValidate(f *testing.F) {
	f.Add("default", "cluster.local")
	f.Fuzz(func(t *testing.T, namespace, dnsDomain string) {
		cfg := &config.Server{
			EnvoyGateway: &egv1a1.EnvoyGateway{
				EnvoyGatewaySpec: egv1a1.EnvoyGatewaySpec{
					Gateway:  egv1a1.DefaultGateway(),
					Provider: egv1a1.DefaultEnvoyGatewayProvider(),
				},
			},
			Namespace: namespace,
			DNSDomain: dnsDomain,
			Logger:    logging.DefaultLogger(egv1a1.LogLevelInfo),
			Elected:   make(chan struct{}),
		}

		err := cfg.Validate()
		if namespace == "" {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	})
}

// Todo: FuzDecode
