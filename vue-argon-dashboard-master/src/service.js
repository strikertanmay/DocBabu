import http from "./api";

class EmployeeDataService {
    getEmployee(name){
        return http.get(`/employee/${name}`)
    }
}

export default new EmployeeDataService();