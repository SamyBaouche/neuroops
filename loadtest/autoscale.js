// CPU pressure on /load (each call burns ~5s of CPU) to drive HPA scale-up.
// Watch it with: kubectl get hpa -n neuroops -w
//   k6 run -e BASE_URL=http://127.0.0.1:PORT loadtest/autoscale.js
import http from "k6/http";
import { check } from "k6";

const BASE_URL = __ENV.BASE_URL || "http://127.0.0.1:8080";

export const options = {
  scenarios: {
    cpu_pressure: {
      executor: "ramping-vus",
      startVUs: 0,
      stages: [
        { duration: "1m", target: Number(__ENV.VUS || 100) },
        { duration: __ENV.HOLD || "5m", target: Number(__ENV.VUS || 100) },
        { duration: "1m", target: 0 },
      ],
    },
  },
};

export default function () {
  const res = http.get(`${BASE_URL}/load`, { timeout: "30s" });
  check(res, { "status is 200": (r) => r.status === 200 });
}
