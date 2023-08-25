package gateway

import (
	"sync"
	"testing"

	"github.com/WuKongIM/WuKongIM/pkg/wkserver"
	"github.com/WuKongIM/WuKongIM/pkg/wkserver/client"
	"github.com/stretchr/testify/assert"
)

func TestDataPipeline(t *testing.T) {

	s := wkserver.New("tcp://127.0.0.1:0")
	err := s.Start()
	assert.NoError(t, err)
	defer func() {
		err = s.Stop()
		assert.NoError(t, err)
	}()

	content := []byte("test")

	var w sync.WaitGroup
	w.Add(len(content))
	s.Route("/node/conn/write", func(c *wkserver.Context) {

		c.WriteOk()
		w.Done()

	})

	cli := client.New(s.Addr().String())
	err = cli.Connect()
	assert.NoError(t, err)

	defer cli.Close()

	d := newDataPipeline(1, "1", cli)
	d.start()
	defer d.stop()

	_, err = d.append([]byte(content))
	assert.NoError(t, err)

	w.Wait()

}
