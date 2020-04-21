package raft

type node struct {
	connect bool
	address string
}

// State def
type State int

// status of node
const (
	Follower State = iota + 1
	Candidate
	Leader
)

// LogEntry struct
type LogEntry struct {
	LogTerm  int
	LogIndex int
	LogCMD   interface{}
}

// Raft Node
type Raft struct {
	me int

	nodes map[int]*node

	state       State
	currentTerm int
	votedFor    int
	voteCount   int

	// 日志条目集合
	log []LogEntry

	// 被提交的最大索引
	commitIndex int
	// 被应用到状态机的最大索引
	lastApplied int

	// 保存需要发送给每个节点的下一个条目索引
	nextIndex []int
	// 保存已经复制给每个节点日志的最高索引
	matchIndex []int

	// channels
	heartbeatC chan bool
	toLeaderC  chan bool
}

// 请求投票
type VoteArgs struct {
	Term        int
	CandidateID int
}

type VoteReply struct {
	//当前任期号， 以便候选人去更新自己的任期号
	Term int
	//候选人赢得此张选票时为真
	VoteGranted bool
}

type HeartbeatArgs struct {
	Term     int
	LeaderID int

	// 新日志之前的索引
	PrevLogIndex int
	// PrevLogIndex 的任期号
	PrevLogTerm int
	// 准备存储的日志条目（表示心跳时为空）
	Entries []LogEntry
	// Leader 已经commit的索引值
	LeaderCommit int
}

type HeartbeatReply struct {
	Success bool
	Term    int

	// 如果 Follower Index小于 Leader Index， 会告诉 Leader 下次开始发送的索引位置
	NextIndex int
}
