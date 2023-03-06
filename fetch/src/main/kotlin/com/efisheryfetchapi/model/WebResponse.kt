package com.efisheryfetchapi.model

import java.util.*

data class WebResponse<T>(
        val message : String,

        val data : T,
)
