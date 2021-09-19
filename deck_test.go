package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
    for _, length := range []int{5, 10, 30, 40} {
        deck := NewDeck(length)
        assert.Equal(t, length, len(deck), "")
    }
}

func TestShuffle(t *testing.T) {
    for _, length := range []int{5, 10, 30, 40} {
        deck := NewDeck(length)
        deck.shuffle()
        assert.Equal(t, length, len(deck), "")
    }
}

func TestPop(t *testing.T) {
    for _, length := range []int{5, 10, 30, 40} {
        deck := NewDeck(length)
        lastCard := deck[len(deck)-1]
        card, deck := pop(deck)

        assert.Equal(t, length-1, len(deck), "")
        assert.Equal(t, lastCard, card, "")
    }    
}

func TestDeckPop(t *testing.T) {
    for _, length := range []int{5, 10, 30, 40} {
        deck := NewDeck(length)
        lastCard := deck[len(deck)-1]
        card := deck.pop()

        assert.Equal(t, length-1, len(deck), "")
        assert.Equal(t, lastCard, card, "")
    }
}

// ------------- POP  --------------------------

/*
type MyInt int
type MyArray []MyInt

func popX(pArray *[]int) int {
    array := *pArray
    *pArray = array[:len(array)-1]
    return array[len(array)-1]
}

func (pArray *MyArray) popY() MyInt {
    array := *pArray
    *pArray = array[:len(array)-1]
    return array[len(array)-1]
}

func pop1(array []int) (int, []int) {
    return array[len(array)-1], array[:len(array)-1]
}

func (array MyArray) pop2() (MyInt, MyArray) {
    return array[len(array)-1], array[:len(array)-1]
}

func (array *MyArray) pop3() MyInt {
    el := (*array)[len((*array))-1]
    (*array) = (*array)[:len((*array))-1]
    return el 
}

func TestPopX(t *testing.T) {
    array := []int{1,2,3,4}
    el := popX(&array)
    //fmt.Println(el, array)

    assert.Equal(t, 4, el, "")
    assert.Equal(t, 3, len(array), "")
    assert.Equal(t, 3, array[len(array)-1], "")
}

func TestPopY(t *testing.T) {
    array := MyArray([]MyInt{MyInt(1),MyInt(2),MyInt(3),MyInt(4)})
    el := array.popY()
    //fmt.Println(el, array)

    assert.Equal(t, MyInt(4), el, "")
    assert.Equal(t, 3, len(array), "The array should be one element shorter.")
    assert.Equal(t, MyInt(3), array[len(array)-1], 
        "The last element should be ...")
}

func TestPop1(t *testing.T) {
    array := []int{1,2,3,4}
    el, array := pop1(array)
    //fmt.Println(el, array)

    assert.Equal(t, 4, el, "")
    assert.Equal(t, 3, len(array), "")
    assert.Equal(t, 3, array[len(array)-1], "")
}

func TestPop2(t *testing.T) {
    array := MyArray([]MyInt{MyInt(1),MyInt(2),MyInt(3),MyInt(4)})
    el, array := array.pop2()
    //fmt.Println(el, array)

    assert.Equal(t, MyInt(4), el, "")
    assert.Equal(t, 3, len(array), "")
    assert.Equal(t, MyInt(3), array[len(array)-1], "")
}

func TestPop3(t *testing.T) {
    array := MyArray([]MyInt{MyInt(1),MyInt(2),MyInt(3),MyInt(4)})
    el := array.pop3()
    //fmt.Println(el, array)

    assert.Equal(t, MyInt(4), el, "")
    assert.Equal(t, 3, len(array), "The array should be one element shorter.")
    assert.Equal(t, MyInt(3), array[len(array)-1], 
        "The last element should be ...")
}
*/

