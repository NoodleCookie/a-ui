package ui

type process string

const (
	Block = process("block")
	Awake = process("awake")
	Stop  = process("stop")
)

type processMode string

const (
	Sync  = processMode("sync,default")
	Async = processMode("async")

	// Mix mode will block your main thread and wait for process Awake | Stop to release
	Mix = processMode("mix")
)
