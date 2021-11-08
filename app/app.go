package app

import (
	"strconv"
	"sync"

	"github.com/barbucatalinn/collatz-conjecture/calculator"
	"github.com/barbucatalinn/collatz-conjecture/fileloader"
)

// App is the main struct of the application
type App struct {
	data []uint64
}

// New creates a new app
func New(filename string) (*App, error) {
	app := new(App)

	// load data
	err := app.loadData(filename)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// loadData loads the csv data
func (app *App) loadData(filename string) error {
	// load CSV data
	csvData, err := fileloader.ReadCSVFile(filename)
	if err != nil {
		return err
	}

	// format to uint64
	for _, i := range csvData {
		if f, err := strconv.ParseUint(i[0], 10, 64); err == nil {
			app.data = append(app.data, f)
		}
	}

	return nil
}

// CalculateCollatzConjectureSteps calls the 'calculator.CalculateCollatzConjectureSteps' function
// and returns a map of the results
func (app *App) CalculateCollatzConjectureSteps() map[uint64]int {
	r := make(map[uint64]int)

	var wg sync.WaitGroup
	ch := make(chan map[uint64]int, len(app.data))

	for i := 0; i < len(app.data); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			calculator.CalculateCollatzConjectureSteps(app.data[i], ch)
		}(i)
	}

	wg.Wait()
	close(ch)

	for m := range ch {
		for k, v := range m {
			r[k] = v
		}
	}

	return r
}
