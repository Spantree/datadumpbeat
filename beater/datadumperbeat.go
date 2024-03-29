package beater

import (
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/manveru/faker"
	"github.com/spantree/datadumpbeat/config"
	"time"
	"math/rand"
	"github.com/dgryski/go-discreterand"
)

var (
	verbAliasTable = discreterand.NewAlias([]float64{0.6, 0.1, 0.1, 0.2}, rand.NewSource(time.Now().UnixNano()))
	httpVerbs      = []string{"GET", "POST", "PATCH", "PUT"}

	respCodeAliasTable = discreterand.NewAlias([]float64{0.9, 0.04, 0.02, 0.04}, rand.NewSource(time.Now().UnixNano()))
	respCodes          = []int{200, 404, 500, 301}

	useragents = []string{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1944.0 Safari/537.36", "Mozilla/5.0 (Linux; U; Android 2.3.5; en-us; HTC Vision Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1", "Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25", "Mozilla/5.0 (Windows; U; Windows NT 6.1; rv:2.2) Gecko/20110201", "Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0", "Mozilla/5.0 (Windows; U; MSIE 9.0; WIndows NT 9.0; en-US))"}
)

type datadumpbeat struct {
	done   chan struct{}
	config config.Config
	faker  *faker.Faker
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	faker, fakerErr := faker.New(config.Locale)
	if fakerErr != nil {
		return nil, fmt.Errorf("Error creating a faker instance with locale: %s", config.Locale, fakerErr)
	}

	bt := &datadumpbeat{
		done:   make(chan struct{}),
		faker:  faker,
		config: config,
	}
	return bt, nil
}

//TODO: Add in random seeding of URL/Referrer
//TODO: Preseed some referrers/pages so traffic looks more uniform
//TODO: Byte code randomization
func (bt *datadumpbeat) Run(b *beat.Beat) error {
	logp.Info("datadumpbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		events := make([]common.MapStr, bt.config.EventsPerPeriod)
		for i := 0; i < bt.config.EventsPerPeriod; i++ {
			events[i] = common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       b.Name,
				"message":    bt.generateFakeLogLine(),
			}
		}
		bt.client.PublishEvents(events)
		logp.Info("Events sent")
	}
}

func (bt *datadumpbeat) generateFakeLogLine() string {
	ip := bt.faker.IPv4Address()
	uri := bt.faker.URL()
	referer := bt.faker.URL()
	time := time.Now().Format("02/Jan/2006:15:04:05 -0700")
	bytes := 123123
	useragent := userAgent()

	return fmt.Sprintf(`%s - - [%s] "%s %s HTTP/1.0" %d %d "%s" "%s"`, ip, time, httpVerb(), uri, respCode(), bytes, referer, useragent)
}

func (bt *datadumpbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func respCode() int {
	return respCodes[respCodeAliasTable.Next()]
}

// Returns a random HTTP verb
func httpVerb() string {
	return httpVerbs[verbAliasTable.Next()]
}

func userAgent() string {
	return useragents[rand.Intn(len(useragents))]
}
