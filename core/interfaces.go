/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package core

type Namespace string

// This
type Wrapper interface {
	GetEventChannel() chan Event
}

type Implementor interface {
}

type Event interface {
}

type Response interface {
}
