package com.efisheryfetchapi.model

import java.util.*

data class AggregateResponse(
        val count : Int,

        val minPrice : Int,
        val minSize : Int,

        val maxPrice : Int,
        val maxSize : Int,

        val avgPrice : Int,
        val avgSize : Int,
)
