package avro_test

import "github.com/electric-saw/avro-codec/v2"

func ConfigTeardown() {
	// Reset the caches
	avro.DefaultConfig = avro.Config{}.Freeze()
}
