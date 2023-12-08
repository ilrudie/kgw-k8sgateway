package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
	vault "github.com/hashicorp/vault/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// NewAuthorizedClient returns a vault client that has been authenticated with the provided settings,
// or an error if construction or authentication fails.
func NewAuthorizedClient(ctx context.Context, vaultSettings *v1.Settings_VaultSecrets) (*vault.Client, error) {
	client, err := newClientForSettings(vaultSettings)
	if err != nil {
		return nil, err
	}

	authMethod, err := newAuthMethodForSettings(ctx, vaultSettings)
	if err != nil {
		return nil, err
	}

	_, err = client.Auth().Login(ctx, authMethod)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func newClientForSettings(vaultSettings *v1.Settings_VaultSecrets) (*vault.Client, error) {
	cfg, err := parseVaultSettings(vaultSettings)
	if err != nil {
		return nil, err
	}

	return api.NewClient(cfg)
}

func parseVaultSettings(vaultSettings *v1.Settings_VaultSecrets) (*api.Config, error) {
	cfg := api.DefaultConfig()

	if addr := vaultSettings.GetAddress(); addr != "" {
		cfg.Address = addr
	}
	if tlsConfig := parseTlsSettings(vaultSettings); tlsConfig != nil {
		if err := cfg.ConfigureTLS(tlsConfig); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func parseTlsSettings(vaultSettings *v1.Settings_VaultSecrets) *api.TLSConfig {
	var tlsConfig *api.TLSConfig

	// helper functions to avoid repeated nilchecking
	addStringSetting := func(s string, addSettingFunc func(string)) {
		if s == "" {
			return
		}
		if tlsConfig == nil {
			tlsConfig = &api.TLSConfig{}
		}
		addSettingFunc(s)
	}
	addBoolSetting := func(b *wrapperspb.BoolValue, addSettingFunc func(bool)) {
		if b == nil {
			return
		}
		if tlsConfig == nil {
			tlsConfig = &api.TLSConfig{}
		}
		addSettingFunc(b.GetValue())
	}

	setCaCert := func(s string) { tlsConfig.CACert = s }
	setCaPath := func(s string) { tlsConfig.CAPath = s }
	setClientCert := func(s string) { tlsConfig.ClientCert = s }
	setClientKey := func(s string) { tlsConfig.ClientKey = s }
	setTlsServerName := func(s string) { tlsConfig.TLSServerName = s }
	setInsecure := func(b bool) { tlsConfig.Insecure = b }

	// Add our settings to the vault TLS config, preferring settings set in the
	// new TlsConfig field if it is used to those in the deprecated fields
	if tlsSettings := vaultSettings.GetTlsConfig(); tlsSettings == nil {
		addStringSetting(vaultSettings.GetCaCert(), setCaCert)
		addStringSetting(vaultSettings.GetCaPath(), setCaPath)
		addStringSetting(vaultSettings.GetClientCert(), setClientCert)
		addStringSetting(vaultSettings.GetClientKey(), setClientKey)
		addStringSetting(vaultSettings.GetTlsServerName(), setTlsServerName)
		addBoolSetting(vaultSettings.GetInsecure(), setInsecure)
	} else {
		addStringSetting(vaultSettings.GetTlsConfig().GetCaCert(), setCaCert)
		addStringSetting(vaultSettings.GetTlsConfig().GetCaPath(), setCaPath)
		addStringSetting(vaultSettings.GetTlsConfig().GetClientCert(), setClientCert)
		addStringSetting(vaultSettings.GetTlsConfig().GetClientKey(), setClientKey)
		addStringSetting(vaultSettings.GetTlsConfig().GetTlsServerName(), setTlsServerName)
		addBoolSetting(vaultSettings.GetTlsConfig().GetInsecure(), setInsecure)
	}

	return tlsConfig

}