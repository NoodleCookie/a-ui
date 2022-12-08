package ui

import (
	"fmt"
)

type BusinessSelectNode struct {
	SelectNode

	mode     processMode
	process  chan process
	business func() error
}

func (b *BusinessSelectNode) Process() chan process {
	return b.process
}

func (b *BusinessSelectNode) Business(business func() error) *BusinessSelectNode {
	b.process = make(chan process, 0)
	b.business = business
	return b
}

func (b *BusinessSelectNode) Mode(mode processMode) *BusinessSelectNode {
	b.mode = mode
	return b
}

func (b *BusinessSelectNode) Do() {
	b.ui.Items = *b.bindData
	_, key, err := b.ui.Run()
	if err != nil {
		b.errCatch(err)
	}

	switch b.mode {
	case Async:
		go func() {
			if err := b.callback(key); err != nil {
				b.errCatch(err)
			}
		}()
		if err := b.business(); err != nil {
			b.errCatch(err)
		}
	case Mix:
		go func() {
			if err := b.callback(key); err != nil {
				b.errCatch(err)
			}
			b.process <- Stop
		}()
		for {
			quitCycle := false
			select {
			case p := <-b.process:
				switch p {
				case Awake:
					if err := b.business(); err != nil {
						b.errCatch(err)
					}
					b.process <- Block
				case Stop:
					quitCycle = true
				default:
					b.errCatch(fmt.Errorf("[INTERNAL_ERROR] Unknown Process: %s", p))
				}
			}
			if quitCycle {
				break
			}
		}
	default:
		if err := b.callback(key); err != nil {
			b.errCatch(err)
		}
		if err := b.business(); err != nil {
			b.errCatch(err)
		}
	}
}
