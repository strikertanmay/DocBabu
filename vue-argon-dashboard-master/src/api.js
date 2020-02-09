import axios from 'axios'

const client = axios.create({
    baseURL: 'http://d94bef41.ngrok.io',
    json: true
  })

  export default {
    async execute(method, resource, data) {
        return client({
            method,
            url: resource,
            data,
            headers: {}
        }).then((req) => {
            return req.data;
        }).catch((err) => {
            console.log(`Caught error invoking API using method of ${method} on resource of ${resource}.`, err);
        });
    },
    createEmployee(data) {
        return this.execute('post', '/xyz', data);
    },
    getEmployee(data) {
        return this.execute('get', '/abc', data);
    },
};