// +build bar,!foo

package keynuker_go_common

import "log"

func init() {
	log.Printf("Setting UseDockerSkeleton = false")
	UseDockerSkeleton = false
}
