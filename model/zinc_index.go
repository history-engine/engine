package model

// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/template.go
// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/index.go

type Index struct {
	ShardNum    int64                  `json:"shard_num"`
	Name        string                 `json:"name"`
	StorageType string                 `json:"storage_type"`
	Settings    *IndexSettings         `json:"settings,omitempty"`
	Mappings    *Mappings              `json:"mappings,omitempty"`
	Shards      map[string]*IndexShard `json:"shards"`
	Stats       IndexStat              `json:"stats"`
	Version     string                 `json:"version"`
}

type IndexShard struct {
	ShardNum int64               `json:"shard_num"`
	ID       string              `json:"id"`
	NodeID   string              `json:"node_id"` // remote instance ID
	Shards   []*IndexSecondShard `json:"shards"`
	Stats    IndexStat           `json:"stats"`
}

type IndexSecondShard struct {
	ID    int64     `json:"id"`
	Stats IndexStat `json:"stats"`
}

type IndexStat struct {
	DocTimeMin  int64  `json:"doc_time_min"`
	DocTimeMax  int64  `json:"doc_time_max"`
	DocNum      uint64 `json:"doc_num"`
	StorageSize uint64 `json:"storage_size"`
	WALSize     uint64 `json:"wal_size"`
}

type IndexSimple struct {
	Name        string                 `json:"name"`
	StorageType string                 `json:"storage_type"`
	ShardNum    int64                  `json:"shard_num"`
	Settings    *IndexSettings         `json:"settings,omitempty"`
	Mappings    map[string]interface{} `json:"mappings,omitempty"`
}

type IndexSettings struct {
	NumberOfShards   int64          `json:"number_of_shards,omitempty"`
	NumberOfReplicas int64          `json:"number_of_replicas,omitempty"`
	Analysis         *IndexAnalysis `json:"analysis,omitempty"`
}

type IndexAnalysis struct {
	Analyzer    map[string]*Analyzer   `json:"analyzer,omitempty"`
	CharFilter  map[string]interface{} `json:"char_filter,omitempty"`
	Tokenizer   map[string]interface{} `json:"tokenizer,omitempty"`
	TokenFilter map[string]interface{} `json:"token_filter,omitempty"`
	Filter      map[string]interface{} `json:"filter,omitempty"` // compatibility with es, alias for TokenFilter
}
