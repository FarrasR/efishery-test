package com.efisheryfetchapi.entity

import java.util.*
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table


@Entity
@Table(name = "products")
data class Product(

        @Id
        val uuid : String,

        @Column(name = "komoditas")
        val komoditas : String,

        @Column(name = "area_provinsi")
        val provinsi : String?,

        @Column(name = "area_kota")
        val kota : String?,

        @Column(name = "size")
        val size : Int,

        @Column(name = "price")
        val price : Int,

        @Column(name = "tglParsed")
        val tglparsed : String,

        @Column(name = "timestamp")
        val timestamp : Int,
)
