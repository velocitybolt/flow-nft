// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Based on https://github.com/wk8/go-ordered-map, Copyright Jean Rougé
 *
 */

package sema

import "container/list"

// MemberStructOrderedMap
//
type MemberStructOrderedMap struct {
	pairs map[*Member]*MemberStructPair
	list  *list.List
}

// NewMemberStructOrderedMap creates a new MemberStructOrderedMap.
func NewMemberStructOrderedMap() *MemberStructOrderedMap {
	return &MemberStructOrderedMap{
		pairs: make(map[*Member]*MemberStructPair),
		list:  list.New(),
	}
}

// Clear removes all entries from this ordered map.
func (om *MemberStructOrderedMap) Clear() {
	om.list.Init()
	// NOTE: Range over map is safe, as it is only used to delete entries
	for key := range om.pairs { //nolint:maprangecheck
		delete(om.pairs, key)
	}
}

// Get returns the value associated with the given key.
// Returns nil if not found.
// The second return value indicates if the key is present in the map.
func (om *MemberStructOrderedMap) Get(key *Member) (result struct{}, present bool) {
	var pair *MemberStructPair
	if pair, present = om.pairs[key]; present {
		return pair.Value, present
	}
	return
}

// GetPair returns the key-value pair associated with the given key.
// Returns nil if not found.
func (om *MemberStructOrderedMap) GetPair(key *Member) *MemberStructPair {
	return om.pairs[key]
}

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *MemberStructOrderedMap) Set(key *Member, value struct{}) (oldValue struct{}, present bool) {
	var pair *MemberStructPair
	if pair, present = om.pairs[key]; present {
		oldValue = pair.Value
		pair.Value = value
		return
	}

	pair = &MemberStructPair{
		Key:   key,
		Value: value,
	}
	pair.element = om.list.PushBack(pair)
	om.pairs[key] = pair

	return
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
func (om *MemberStructOrderedMap) Delete(key *Member) (oldValue struct{}, present bool) {
	var pair *MemberStructPair
	pair, present = om.pairs[key]
	if !present {
		return
	}

	om.list.Remove(pair.element)
	delete(om.pairs, key)
	oldValue = pair.Value

	return
}

// Len returns the length of the ordered map.
func (om *MemberStructOrderedMap) Len() int {
	return len(om.pairs)
}

// Oldest returns a pointer to the oldest pair.
func (om *MemberStructOrderedMap) Oldest() *MemberStructPair {
	return listElementToMemberStructPair(om.list.Front())
}

// Newest returns a pointer to the newest pair.
func (om *MemberStructOrderedMap) Newest() *MemberStructPair {
	return listElementToMemberStructPair(om.list.Back())
}

// Foreach iterates over the entries of the map in the insertion order, and invokes
// the provided function for each key-value pair.
func (om *MemberStructOrderedMap) Foreach(f func(key *Member, value struct{})) {
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		f(pair.Key, pair.Value)
	}
}

// MemberStructPair
//
type MemberStructPair struct {
	Key   *Member
	Value struct{}

	element *list.Element
}

// Next returns a pointer to the next pair.
func (p *MemberStructPair) Next() *MemberStructPair {
	return listElementToMemberStructPair(p.element.Next())
}

// Prev returns a pointer to the previous pair.
func (p *MemberStructPair) Prev() *MemberStructPair {
	return listElementToMemberStructPair(p.element.Prev())
}

func listElementToMemberStructPair(element *list.Element) *MemberStructPair {
	if element == nil {
		return nil
	}
	return element.Value.(*MemberStructPair)
}