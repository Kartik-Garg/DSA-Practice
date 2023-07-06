package com.example;

public class NumberExampleRecursion {
    public static void main(String[] args) {
        // Let's print till 5 in recursion.
        print(1);
    }

    public static void print(int i) {
        // 1. Base condition to stop calling it again and again
        if (i > 5) {
            return;
        }
        // 2. Every call to the same function can be treated as a different function.
        // Since, it will be stored separately in the func calling stack.
        System.out.println(i);
        print(i + 1);
    }
}
