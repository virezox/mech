package main
import "github.com/refraction-networking/utls"

var ciphers = []uint16{
   tls.TLS_AES_256_GCM_SHA384,
   tls.TLS_CHACHA20_POLY1305_SHA256,
   tls.TLS_AES_128_GCM_SHA256,
   tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
   tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
   tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
   tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
   tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
   tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
   tls.FAKE_TLS_DHE_RSA_WITH_AES_256_GCM_SHA384,
   tls.FAKE_TLS_DHE_RSA_WITH_AES_128_GCM_SHA256,
   0xCCAA, // TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256
   0xC0AF, // TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8
   0xC0AD, // TLS_ECDHE_ECDSA_WITH_AES_256_CCM
   0xC0AE, // TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8
   0xC0AC, // TLS_ECDHE_ECDSA_WITH_AES_128_CCM
   tls.DISABLED_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,
   tls.DISABLED_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,
   tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
   tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
   tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
   tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
   0xC0A3, // TLS_DHE_RSA_WITH_AES_256_CCM_8
   0xC09F, // TLS_DHE_RSA_WITH_AES_256_CCM
   0xC0A2, // TLS_DHE_RSA_WITH_AES_128_CCM_8
   0xC09E, // TLS_DHE_RSA_WITH_AES_128_CCM
   0x006B, // TLS_DHE_RSA_WITH_AES_256_CBC_SHA256
   0x0067, // TLS_DHE_RSA_WITH_AES_128_CBC_SHA256
   tls.FAKE_TLS_DHE_RSA_WITH_AES_256_CBC_SHA,
   tls.FAKE_TLS_DHE_RSA_WITH_AES_128_CBC_SHA,
   tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
   tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
   0xC0A1, // TLS_RSA_WITH_AES_256_CCM_8
   0xC09D, // TLS_RSA_WITH_AES_256_CCM
   0xC0A0, // TLS_RSA_WITH_AES_128_CCM_8
   0xC09C, // TLS_RSA_WITH_AES_128_CCM
   tls.DISABLED_TLS_RSA_WITH_AES_256_CBC_SHA256,
   tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
   tls.TLS_RSA_WITH_AES_256_CBC_SHA,
   tls.TLS_RSA_WITH_AES_128_CBC_SHA,
   tls.FAKE_TLS_EMPTY_RENEGOTIATION_INFO_SCSV,
}

var preset = &tls.ClientHelloSpec{
   CipherSuites: ciphers,
   Extensions: []tls.TLSExtension{
      &tls.SNIExtension{},
      &tls.SupportedPointsExtension{
         SupportedPoints: []byte{0, 1, 2},
      },
      &tls.SupportedCurvesExtension{
         []tls.CurveID{
            tls.X25519,
            tls.CurveP256,
            0x001E, // Curve448
            tls.CurveP521,
            tls.CurveP384,
         },
      },
      &tls.SessionTicketExtension{},
      &tls.GenericExtension{Id:0x0016}, // encrypt_then_mac
      &tls.UtlsExtendedMasterSecretExtension{},
      &tls.SignatureAlgorithmsExtension{
         SupportedSignatureAlgorithms: []tls.SignatureScheme{
            tls.ECDSAWithP256AndSHA256,
            tls.PKCS1WithSHA256,
         },
      },
      &tls.SupportedVersionsExtension{
         []uint16{tls.VersionTLS12},
      },
      &tls.PSKKeyExchangeModesExtension{
         []uint8{tls.PskModeDHE},
      },
      &tls.KeyShareExtension{
         []tls.KeyShare{
            {Group: tls.X25519},
         },
      },
   },
}
