resource "routeros_queue_type" "test" {
  name           = "pcq-test"
  kind           = "pcq"
  pcq_rate       = 0
  pcq_limit      = 50
  pcq_classifier = ["dst-address"]
}
