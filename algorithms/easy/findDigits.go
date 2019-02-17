package main

func findDigits(n int32) int32 {
    var count int32
    copy := n
    for copy > 0 {
        if copy == 1 {
            count++
            break
        }
        reminder := copy % 10
        if reminder != 0 && n % reminder == 0 {
            count++
            copy -= reminder
        }
        copy /= 10
    }
        return count

}
