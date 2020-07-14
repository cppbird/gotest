package tree

import "sync"

// TrieNode .
type TrieNode struct {
	children map[rune]*TrieNode
}

// Trie .
type Trie struct {
	mu   sync.Mutex
	root *TrieNode
	size int
}
