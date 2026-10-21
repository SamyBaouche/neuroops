// Steady user traffic on / to measure throughput, latency, and availability.
//   k6 run -e BASE_URL=http://127.0.0.1:PORT loadtest/traffic.js
//   k6 run -e BASE_URL=... -e VUS=100 -e DURATION=10m --summary-export=loadtest/results/traffic.json loadtest/traffic.js
import http from "k6/http";
import { check } from "k6";

const BASE_URL = __ENV.BASE_URL || "http://127.0.0.1:8080";

export const options = {
  vus: Number(__ENV.VUS || 100),
  duration: __ENV.DURATION || "2m",
  summaryTrendStats: ["avg", "p(50)", "p(95)", "p(99)", "max"],
  thresholds: {
    // Mirrors the 99.5% availability SLO and the 200ms p99 latency alert.
    http_req_failed: ["rate<0.005"],
    "http_req_duration{expected_response:true}": ["p(99)<200"],
  },
};

export default function () {
  const res = http.get(`${BASE_URL}/`);
  check(res, { "status is 200": (r) => r.status === 200 });
}
