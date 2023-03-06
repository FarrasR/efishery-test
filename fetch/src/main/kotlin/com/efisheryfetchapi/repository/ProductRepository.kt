package com.efisheryfetchapi.repository

import com.efisheryfetchapi.entity.Product
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.jpa.repository.Query
import org.springframework.data.repository.query.Param
import java.util.*

interface ProductRepository : JpaRepository<Product, String>{
    fun findByUuid(uuid: String): Product

//    @Query(value = "from Products p where area_provinsi = #{areaProvinsi} AND tgl_parsed BETWEEN #{startDate} AND #{endDate}")
//    fun findAllBetweenX(@Param("areaProvinsi") areaProvinsi: String ,@Param("startDate") startDate: Date ,@Param("endDate") endDate: Date) : List<Product>


    fun findAllByProvinsiAndTglparsedBetween(provinsi: String, to: String, from: String): List<Product>

    fun findAllByTglparsedBetween(to: String, from: String): List<Product>

    fun findAllByProvinsi(provinsi: String): List<Product>

    fun findAllByTglparsedGreaterThanEqualAndTglparsedLessThanEqual(to: Date, from: Date): List<Product>

}