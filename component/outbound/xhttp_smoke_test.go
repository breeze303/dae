/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2022-2025, daeuniverse Organization <dae@v2raya.org>
 */

package outbound

import (
	"strings"
	"testing"

	"github.com/daeuniverse/dae/component/outbound/dialer"
	"github.com/daeuniverse/dae/config"
	"github.com/sirupsen/logrus"
)

func TestVLESSXHTTPLinkConstructsThroughOutboundNewFromLink(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.PanicLevel)

	option := dialer.NewGlobalOption(&config.Global{}, logger)
	iOption := dialer.InstanceOption{DisableCheck: true}
	link := "vless://00000000-0000-0000-0000-000000000000@example.com:443?type=xhttp&security=none&path=/api&mode=packet-up#xhttp"

	d, err := dialer.NewFromLink(option, iOption, link, "dae-xhttp-smoke")
	if err != nil {
		t.Fatalf("dialer.NewFromLink() error = %v", err)
	}
	if d == nil {
		t.Fatal("dialer.NewFromLink() returned nil dialer")
	}

	prop := d.Property()
	if prop == nil {
		t.Fatal("dialer property is nil")
	}
	if prop.Protocol != "vless" {
		t.Fatalf("Property.Protocol = %q, want vless", prop.Protocol)
	}
	if prop.Address != "example.com:443" {
		t.Fatalf("Property.Address = %q, want example.com:443", prop.Address)
	}
	if prop.SubscriptionTag != "dae-xhttp-smoke" {
		t.Fatalf("Property.SubscriptionTag = %q, want dae-xhttp-smoke", prop.SubscriptionTag)
	}
	if prop.Link == "" {
		t.Fatal("Property.Link is empty")
	}
	if strings.HasPrefix(prop.Link, "xhttp://") {
		t.Fatalf("Property.Link = %q, must stay on vless:// export path", prop.Link)
	}
	if !strings.HasPrefix(prop.Link, "vless://") {
		t.Fatalf("Property.Link = %q, want vless export", prop.Link)
	}
}
