package main

import (
	"log"

	"github.com/yitsushi/go-misskey/services/meta"
)

func getInstanceMeta() meta.InstanceMetaResponse {
	meta, err := client.Meta().InstanceMeta(false)
	if err != nil {
		log.Printf("[meta] Error happened: %s", err)
	}
	return meta
}
