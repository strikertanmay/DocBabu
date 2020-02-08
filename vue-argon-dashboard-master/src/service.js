import http from "./api";

class EmployeeDataService {
    getEmployee(name){
        return http.get(`/employee/${name}`)
    }

    postDocument(data){
        return http.post("/document", data)
    }
}

export default new EmployeeDataService();