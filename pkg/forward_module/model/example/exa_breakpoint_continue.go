package example

import (
	"ldacs_sim_sgw/pkg/forward_module/forward_global"
)

// file struct, 文件结构体
type ExaFile struct {
	forward_global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	forward_global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
