import http from 'k6/http';
import { check, sleep, group } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 300 },
    { duration: '30s', target: 500 },
    { duration: '30s', target: 1000 },
  ],
  thresholds: {
    http_req_duration: ['p(95) < 100'],
    http_req_failed: ['rate < 0.01'],
  },
};

export default function () {
  group('List SPU', () => {
    const url = 'http://localhost:8080/product/spu/list';
        const res = http.get(url);
    check(res, {
      'status is 200': (r) => r.status === 200,
      'response has code 200': (r) => {
        try {
          const json = JSON.parse(r.body);
          return json.code === 200;
        } catch (e) {
          return false;
        }
      },
      'response has data array': (r) => {
        try {
          const json = JSON.parse(r.body);
          return Array.isArray(json.data) && json.data.length > 0;
        } catch (e) {
          return false;
        }
      }
    });
    sleep(1);
  });
}