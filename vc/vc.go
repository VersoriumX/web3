package vc

import (
	"time"
)

type VerifiableCredential struct {
	Context           []string               `json:"@context,omitempty"`
	ID                string                 `json:"id,omitempty"`
	Type              []string               `json:"type,omitempty"`
	Issuer            string                 `json:"issuer,omitempty"`
	IssuanceDate      *time.Time             `json:"issuanceDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	Proof             *Proof                 `json:"proof,omitempty"`
}

// NewVerifiableCredential returns a new instance of VerifiableCredential
// with the default context and type assigned.
func NewVerifiableCredential() *VerifiableCredential {
	return &VerifiableCredential{
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		Type:    []string{"VerifiableCredential"},
	}
}

// VerifiablePresentation combines one or more credentials.
//
// https://www.w3.org/TR/verifiable-claims-data-model/#presentations
type VerifiablePresentation struct {
	Context              []string                `json:"@context,omitempty"`
	ID                   string                  `json:"id,omitempty"`
	Type                 []string                `json:"type,omitempty"`
	VerifiableCredential []*VerifiableCredential `json:"verifiableCredential,omitempty"`
	Proof                *Proof                  `json:"proof,omitempty"`
}

// NewVerifiablePresentation returns a new instance of VerifiablePresentation
// with the default context and type assigned.
func NewVerifiablePresentation() *VerifiablePresentation {
	return &VerifiablePresentation{
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		Type:    []string{"VerifiablePresentation"},
	}
}

type Proof struct {
	Type       string     `json:"type,omitempty"`
	Created    *time.Time `json:"created,omitempty"`
	Creator    string     `json:"creator,omitempty"`
	ProofValue string     `json:"proofValue,omitempty"`
}
{
  "digest": "sha256:191cc78a3f3ce46c786c1b94806fcc2ce4c6cd630f4915151cb7f5671c7ed859",
  "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
  "size": 2000,
  "config": {
    "digest": "sha256:c464262ffe6397e69b7521714aaf0af2f6fb0c066b8f319d7bd556665e0eb3a7",
    "mediaType": "application/vnd.docker.container.image.v1+json",
    "size": 8492
  },
  "layers": [
    {
      "digest": "sha256:af107e978371b6cd6339127a05502c5eacd1e6b0e9eb7b2f4aa7b6fc87e2dd81",
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 29125963
    },
    {
      "digest": "sha256:8ce3f2b601ccac03ff1858022363c325355bafba224123a4563dade58bc8e70f",
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 3507550
    },
    {
      "digest": "sha256:2c1fe8bcf114038da9146a2528b2c3b606a09873aee7122f8529d45418fcd9e9",
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 3363523
    },
    {
      "digest": "sha256:64119eae2e5157ca18d0ec9b9c7865e454b09bca702c3e2f579703deb0eaaeaf",
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 12376761
    },
    {
      "digest": "sha256:a28cd3d033268c30cb986f4a0d36c4aeda4c5aeaa5c764562f323c73d1f4f771",
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "size": 244
    },
    {
      "digest": "sha256:c0cfe091195bd65ef88af5bc5eff6d0178f093ec7af0b16a5821ee3c58e7bc8a",
      "mediaType": "application/x-gzip",
      "size": 93
    },
    {
      "digest": "sha256:cd5ad8c369bea6d3363bbef161f362ba16b9d3a30c8532660255a69276b61f49",
      "mediaType": "application/x-gzip",
      "size": 7290818
    },
    {
      "digest": "sha256:41e9712bd159d653f0c0a091b2f249412fb12cd5f944d1328b13755fc7dac324",
      "mediaType": "application/x-gzip",
      "size": 236472402
    },
    {
      "digest": "sha256:c464262ffe6397e69b7521714aaf0af2f6fb0c066b8f319d7bd556665e0eb3a7",
      "mediaType": "text/plain; charset=utf-8",
      "size": 8492
    }
  ],
  "labels": {
    "org.opencontainers.image.url": "https://github.com/VersoriumX/LabVX",
    "org.opencontainers.image.licenses": "Apache-2.0",
    "org.opencontainers.image.title": "LabVX",
    "org.opencontainers.image.revision": "a2cf3974f4cc7aca0198883efab307c546f36fca",
    "org.opencontainers.image.created": "2024-01-07T08:08:48.609Z",
    "org.opencontainers.image.version": "edge",
    "org.opencontainers.image.description": "Analysis & simulation repo for Delv",
    "org.opencontainers.image.source": "https://github.com/VersoriumX/LabVX"
  }
}
