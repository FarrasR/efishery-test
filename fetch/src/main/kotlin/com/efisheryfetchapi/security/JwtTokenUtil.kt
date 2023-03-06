package com.efisheryfetchapi.security

import org.springframework.boot.autoconfigure.security.oauth2.resource.OAuth2ResourceServerProperties.Jwt
import org.springframework.stereotype.Component
import io.jsonwebtoken.Jwts
import org.springframework.stereotype.Service

@Service
class JwtTokenUtil {
    private val secret = "123456789012345678901234567890123456789012345678901234567890"

    private fun getClaims(token: String) =
        Jwts.parserBuilder().setSigningKey(secret.toByteArray()).build().parseClaimsJws(token)

    fun getUsername(token: String): String = getClaims(token).body.subject

    fun isTokenRoleAdmin(token: String): Boolean {
        try  {
            if (getClaims(token).body["role"].toString() == "admin") {
                return true
            }
            return false
        } catch (e: Exception){
            return false
        }
    }

    fun isTokenValid(token: String): Boolean {
        return try  {
            getClaims(token)
            true
        } catch (e: Exception){
            false
        }
    }
}