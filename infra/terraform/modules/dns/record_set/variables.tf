variable "PROJECT_ID" {
  type = string
}
variable "name" {
  type = string
}
variable "dns_zone_name" {
  type = string
}
variable "ttl" {
  type = number
}
variable "rrdatas" {
  type = list(string)
}

variable "record_type" {
  type = string
  validation {
    error_message = "invalid record_type"
    condition = contains([
      "A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "HTTPS", "IPSECKEY", "MX", "NAPTR", "NS", "PTR", "SPF", "SRV",
      "SSHFP", "SVCB", "TLSA", "TXT", "SOA"
    ], var.record_type)
  }
}
