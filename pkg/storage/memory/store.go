package memory

import "github.com/xephonhq/xephon-k/pkg/common"

// Store is the in memory storage with data and index
type Store struct {
	data  Data
	index *Index // TODO: might change to value instead of pointer
}

// NewMemStore creates an in memory storage with small allocated space
func NewMemStore() *Store {
	store := Store{}
	// TODO: add a function to create the data?
	store.data = make(Data, initSeriesCount)
	store.index = NewIndex(initSeriesCount)
	return &store
}

// StoreType implements Store interface
func (store Store) StoreType() string {
	return "memory"
}

// QueryIntSeriesBatch implements Store interface
func (store Store) QueryIntSeriesBatch(queries []common.Query) ([]common.QueryResult, []common.IntSeries, error) {
	result := make([]common.QueryResult, 0, len(queries))
	series := make([]common.IntSeries, 0, len(queries))
	// TODO:
	// - first look up the series id
	// - add match number
	// - read the data by time range
	// - apply the aggregator when look up?
	// - test it in non e2e test
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		queryResult := common.QueryResult{Query: query, Matched: 0}
		switch query.MatchPolicy {
		case "exact":
			seriesID := common.Hash(&query)
			oneSeries, ok := store.data[seriesID]
			if ok {
				queryResult.Matched = 1
				series = append(series, oneSeries.ReadByStartEndTime(query.StartTime, query.EndTime))
			}
		case "filter":
			// TODO: we should also expose a HTTP API for query series ID only
			// FIXME: this is a dirty hack to be compatible with the Name filed in the query, it is treated as __name__ tag
			// need to make a shallow copy, otherwise it will refer to itself and cause stackoverflow
			originalFilter := query.Filter
			query.Filter = common.Filter{Type: "and", LeftOperand: &common.Filter{Type: "tag_match", Key: nameTagKey, Value: query.Name},
				RightOperand: &originalFilter}
			seriesIDs := store.index.Filter(&query.Filter)
			queryResult.Matched = len(seriesIDs)
			for j := 0; j < len(seriesIDs); j++ {
				// TODO: let's just assume all series in the index is all in the memory, so we don't check the data map
				seriesID := seriesIDs[j]
				series = append(series, store.data[seriesID].ReadByStartEndTime(query.StartTime, query.EndTime))
			}
		default:
			// TODO: query the index to do the filter
			log.Warn("non exact match is not supported!")
		}
		result = append(result, queryResult)
	}
	return result, series, nil
}

// QueryIntSeries implements Store interface
// Deprecated: Use QueryIntSeriesBatch instead
func (store Store) QueryIntSeries(query common.Query) ([]common.IntSeries, error) {
	series := make([]common.IntSeries, 0)
	// TODO: not hard coded string
	switch query.MatchPolicy {
	case "exact":
		// fetch the series
		seriesID := common.Hash(&query)
		// TODO: should we make a copy of the points, what would happen if there are
		// write when we are encoding it to json
		// TODO: there is mutex on IntSeries store, how does prometheus etc. handle this?
		// should we have a get method or things like that?
		// prometheus use Iterator .... maybe we need custom implements, I think it also have blocks
		oneSeries, ok := store.data[seriesID]
		if ok {
			series = append(series, oneSeries.ReadByStartEndTime(query.StartTime, query.EndTime))
		}
		return series, nil
	case "filter":
		// TODO: real filter
		log.Warn("TODO: write code for filter")
	default:
		// TODO: query the index to do the filter
		log.Warn("non exact match is not supported!")
	}
	return series, nil
}

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	// TODO: will using range and array access have difference
	for _, oneSeries := range series {
		id := common.Hash(&oneSeries)
		// TODO: this should return error and we should handle it somehow
		// Write Data
		store.data.WriteIntSeries(id, oneSeries)
		// Write Index
		// NOTE: we store series name as special tag
		store.index.Add(id, nameTagKey, oneSeries.Name)
		for k, v := range oneSeries.Tags {
			store.index.Add(id, k, v)
		}
	}
	return nil
}

// Shutdown TODO: gracefully flush in memory data to disk
func (store Store) Shutdown() {
	log.Info("shutting down memoery store, nothing to do, have a nice weekend~")
}
