package com.efisheryfetchapi

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class FetchApiApplication

fun main(args: Array<String>) {
	runApplication<FetchApiApplication>(*args)
}
