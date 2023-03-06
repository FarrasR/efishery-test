package com.efisheryfetchapi.service

import com.efisheryfetchapi.model.AggregateResponse
import com.efisheryfetchapi.model.ProductResponse
import org.sqlite.Function.Aggregate
import java.util.Date

interface ProductService {
    fun getByUuid(uuid: String): ProductResponse
    fun aggregateProducts(startDate: String, endDate: String, provinsi: String): AggregateResponse
    fun getProducts(startDate: String, endDate: String, provinsi: String): List<ProductResponse>
}