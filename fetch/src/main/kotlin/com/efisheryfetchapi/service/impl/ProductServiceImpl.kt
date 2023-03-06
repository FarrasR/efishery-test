package com.efisheryfetchapi.service.impl

import com.efisheryfetchapi.model.AggregateResponse
import com.efisheryfetchapi.model.ProductResponse
import com.efisheryfetchapi.repository.ProductRepository
import com.efisheryfetchapi.service.ProductService
import org.springframework.stereotype.Service
import java.util.*

@Service
class ProductServiceImpl(val productRepository: ProductRepository,) : ProductService {
    override fun getByUuid(uuid: String): ProductResponse {

        val product = productRepository.findByUuid(uuid)

        return ProductResponse(
            uuid = product.uuid,
            komoditas = product.komoditas,
            areaProvinsi = product.provinsi,
            areaKota = product.kota,
            size = product.size,
            price = product.size,
            tglParsed = product.tglparsed,
            timestamp = product.timestamp,
        )
    }

    override fun aggregateProducts(startDate: String, endDate: String, provinsi: String): AggregateResponse{
        //VERY MESSY ALGORITHM
        val products = productRepository.findAllByProvinsiAndTglparsedBetween(provinsi, startDate, endDate)

        if (products.size == 0) {
            return AggregateResponse(
                count = 0,
                minPrice =0,
                minSize =0,
                maxPrice = 0,
                maxSize = 0,
                avgPrice =0,
                avgSize = 0,
            )
        }

        val prices : Array<Int> = arrayOf()
        val sizes : Array<Int> = arrayOf()

        var totalPrices = 0
        var totalSizes = 0
        var count = 0

        var maxPrice = -100000
        var maxSize = -100000

        var minPrice = 1000000000
        var minSize = 1000000000
        for (product in products) {
            count +=1
            totalPrices += product.price
            totalSizes += product.size
            if (product.size > maxSize) {
                maxSize = product.size
            }
            if (product.price > maxPrice) {
                maxPrice = product.price
            }
            if (product.size < minSize) {
                minSize = product.size
            }
            if (product.price < minPrice) {
                minPrice = product.price
            }
        }

        return AggregateResponse(
            count = count,
            minPrice =minPrice,
            minSize =minSize,
            maxPrice = maxPrice,
            maxSize = maxSize,
            avgPrice =totalPrices/count,
            avgSize = totalSizes/count,
        )
    }

    override fun getProducts(startDate: String, endDate: String, provinsi: String): List<ProductResponse> {
        val products = productRepository.findAllByProvinsiAndTglparsedBetween(provinsi, startDate, endDate)
        var listProductResponse = ArrayList<ProductResponse>()
        for (product in products) {

            val temp = ProductResponse(
                uuid = product.uuid,
                komoditas = product.komoditas,
                areaProvinsi = product.provinsi,
                areaKota = product.kota,
                size = product.size,
                price = product.size,
                tglParsed = product.tglparsed,
                timestamp = product.timestamp,
            )

            listProductResponse.add(temp)
        }
        return listProductResponse
    }
}