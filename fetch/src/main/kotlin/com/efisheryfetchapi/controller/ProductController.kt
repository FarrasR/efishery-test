package com.efisheryfetchapi.controller

import com.efisheryfetchapi.model.ProductResponse
import com.efisheryfetchapi.model.WebResponse
import com.efisheryfetchapi.security.JwtTokenUtil
import com.efisheryfetchapi.service.ProductService
import io.jsonwebtoken.Jwts
import org.springframework.web.bind.annotation.*
import java.text.SimpleDateFormat
import java.util.Date

@RestController
class ProductController(val productService: ProductService,
                        val jwtTokenUtil: JwtTokenUtil) {
    @GetMapping(
            value = ["/products/{uuid}"],
            produces = ["application/json"]
    )
    fun getByUuid(@PathVariable("uuid") uuid: String,
                  @RequestHeader("Authorization") authorization: String): WebResponse<Any?> {
        if (!jwtTokenUtil.isTokenValid(authorization)) {
            return WebResponse(
                message = "unauthorized",
                data = null
            )
        }
        val product =  productService.getByUuid(uuid)
        return WebResponse(
            message = "success",
            data = product
        )
    }

    @GetMapping(
        value = ["/products/"],
        produces = ["application/json"]
    )
    fun getProducts(@RequestParam("provinsi") provinsi: String,
                    @RequestParam("startDate") startDate: String,
                    @RequestParam("endDate") endDate: String,
                    @RequestHeader("Authorization") authorization: String): WebResponse<Any?> {

        if (!jwtTokenUtil.isTokenValid(authorization)) {
            return WebResponse(
                message = "unauthorized",
                data = null
            )
        }
        val products = productService.getProducts(startDate, endDate, provinsi)

        return WebResponse(
            message = "success",
            data = products
        )
    }


    @GetMapping(
        value = ["/products-aggregate/"],
        produces = ["application/json"]
    )
    fun aggregateProducts(@RequestParam("provinsi") provinsi: String,
                          @RequestParam("startDate") startDate: String,
                          @RequestParam("endDate") endDate: String,
                          @RequestHeader("Authorization") authorization: String): WebResponse<Any?> {

        if (!jwtTokenUtil.isTokenRoleAdmin(authorization)) {
            return WebResponse(
                message = "unauthorized",
                data = null
            )
        }
        val aggregate = productService.aggregateProducts(startDate, endDate, provinsi)

        return WebResponse(
            message = "success",
            data = aggregate
        )
    }
}