// +build UseDockerSkeleton

package keynuker_go_common

import "log"

func init() {
	log.Printf("Setting UseDockerSkeleton = true")
	UseDockerSkeleton = true
}
