package training

// 8.6 `M` Design Twitter

/*** @LEETCODE https://leetcode.com/problems/design-twitter/
Design a simplified version of Twitter where users can post tweets,
follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.
Implement the Twitter class:
- Twitter() Initializes your twitter object.
- void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
- List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
- void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
- void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.
---
Example 1:
Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]
Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2);    // User 1 follows user 2.
twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  // User 1 unfollows user 2.
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
---
Constraints:
1 <= userId, followerId, followeeId <= 500
0 <= tweetId <= 10^4
All the tweets have unique IDs.
At most 3 * 10^4 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=pNichitDD2E
***/

type Tweet struct {
	userId  int
	tweetId int
	time    int // Timestamp to maintain chronological order
}

type Twitter struct {
	tweets      []Tweet                // Global tweet feed (chronological order)
	userTweets  map[int][]Tweet        // User-specific tweets
	followGraph map[int]map[int]struct{} // Followers and followees
	time        int                    // Simulate timestamp
}

func Constructor() Twitter {
	return Twitter{
		tweets:      []Tweet{},
		userTweets:  make(map[int][]Tweet),
		followGraph: make(map[int]map[int]struct{}),
		time:        0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++
	tweet := Tweet{userId: userId, tweetId: tweetId, time: this.time}
	this.tweets = append(this.tweets, tweet)
	this.userTweets[userId] = append(this.userTweets[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// Use a max heap to efficiently get the top 10 recent tweets
	h := &TweetHeap{}
	heap.Init(h)

	// Add user's tweets to the heap
	for _, tweet := range this.userTweets[userId] {
		heap.Push(h, tweet)
	}

	// Add followees' tweets to the heap
	if followees, ok := this.followGraph[userId]; ok {
		for followeeId := range followees {
			for _, tweet := range this.userTweets[followeeId] {
				heap.Push(h, tweet)
			}
		}
	}

	// Extract the 10 most recent tweets
	var newsFeed []int
	for i := 0; i < 10 && h.Len() > 0; i++ {
		newsFeed = append(newsFeed, heap.Pop(h).(Tweet).tweetId)
	}
	return newsFeed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followGraph[followerId]; !ok {
		this.followGraph[followerId] = make(map[int]struct{})
	}
	this.followGraph[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.followGraph[followerId]; ok {
		delete(this.followGraph[followerId], followeeId)
	}
}

// TweetHeap is a max-heap for Tweets based on their timestamp
type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	tweet := old[n-1]
	*h = old[:n-1]
	return tweet
}

