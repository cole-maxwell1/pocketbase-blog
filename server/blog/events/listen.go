package events

import "github.com/pocketbase/pocketbase"

func StartEventListens(pb *pocketbase.PocketBase) {
	notifyPosterOnComment(pb)
	// add more event listeners here...

}
