package com.efisheryfetchapi.model

import java.util.*

data class ProductResponse(
        val uuid : String,

        val komoditas : String,

        val areaProvinsi : String?,

        val areaKota : String?,

        val size : Int,

        val price : Int,

        val tglParsed : String,

        val timestamp : Int
)
