package main

func f1(f2 func()) {

}

func f2(f3 func()) {
	f3()
}

func f3() {

}
