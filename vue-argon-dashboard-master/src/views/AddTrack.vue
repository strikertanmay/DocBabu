<template>
  <div>
    <base-header
      class="header pb-8 pt-5 pt-lg-8 d-flex align-items-center"
      style="min-height: 600px; background-size: cover; background-position: center top;"
    >
      <!-- Mask -->
      <span class="mask bg-gradient-success opacity-8"></span>
      <!-- Header container -->
      <div class="container-fluid d-flex align-items-center">
        <div class="row">
          <div class="col-lg-7 col-md-10">
            <h1 class="display-2 text-white">Create A Track</h1>
            <p class="text-white mt-0 mb-5" style="font-size:20px">Create a new track to trace a new file. Add the file's name, the path it will follow for complete processing and track it to completion.The stages of
              processing maybe altered at any possible point by the authorised employees. A QR code for the file will be generated at this stage to facilitate the tracking of the document.
            </p>
          </div>
        </div>
      </div>
    </base-header>

    <div class="container-fluid mt--7">
      <div class="row">
        <div class="col-xl-12 order-xl-1">
          <card shadow type="secondary">
            <div slot="header" class="bg-white border-0">
              <div class="row align-items-center">
                <div class="col-8">
                  <h3 class="mb-0">New Track</h3>
                </div>
                <!-- <div class="col-4 text-right">
                                    <a href="#!" class="btn btn-sm btn-primary">Settings</a>
                </div>-->
              </div>
            </div>
            <template>
              <form @submit.prevent>
                <h6 class="heading-small text-muted mb-4">File Information</h6>
                <div class="pl-lg-4">
                  <div class="row">
                    <div class="col-md-12">
                      <base-input
                        alternative
                        label="File name"
                        placeholder="enter file name"
                        input-classes="form-control-alternative"
                        v-model="model.filename"
                      />
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-md-12">
                      <base-input
                        alternative
                        label="Created By"
                        placeholder="employee name"
                        input-classes="form-control-alternative"
                        v-model="model.creatorName"
                      />
                    </div>
                  </div>
                </div>
                <hr class="my-4" />
                <!-- Address -->
                <h6 class="heading-small text-muted mb-4">Track Information</h6>
                <div class="pl-lg-4">
                  <div class="row" v-for="(input,k) in model.path" :key="k">
                    <div class="col-lg-4">
                      <base-input
                        alternative
                        label="Name"
                        placeholder="name"
                        input-classes="form-control-alternative"
                        v-model="model.path[k].name"
                      />
                    </div>
                    <div class="col-lg-3">
                      <base-input
                        alternative
                        label="Priority"
                        placeholder="priority of the stage"
                        input-classes="form-control-alternative"
                        v-model="model.path[k].priority"
                      />
                    </div>
                    <div class="col-lg-4">
                      <base-input
                        alternative
                        label="Remark"
                        placeholder="remark"
                        input-classes="form-control-alternative"
                        v-model="model.path[k].comment"
                      />
                    </div>
                    <span>
                      <i
                        class="fas fa-plus-circle"
                        @click="add(k)"
                        v-show="k == model.path.length-1"
                      ></i>
                    </span>
                  </div>
                </div>
                <hr class="my-4" />
                <!-- Description -->
                <h6 class="heading-small text-muted mb-4">Additional Information</h6>
                <div class="pl-lg-4">
                  <div class="form-group">
                    <base-input 
                    alternative label="Remarks"
                    v-model="model.path.remarks"
                    >
                      <textarea
                        rows="4"
                        class="form-control form-control-alternative"
                        placeholder="A few words about you ..."
                      >Add some remarks over here.</textarea>
                    </base-input>
                  </div>
                </div>

                <div class="row">
                <div class="pl-lg-4">
                  <base-button type="primary" class="my-4" @click="handleSubmit">Submit</base-button>
                </div>
                </div>
              </form>
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import EmployeeDataService from '@/service';
export default {
  name: "user-profile",
  data() {
    return {
      model: {
        filename: "",
        creatorName: "",
        path: [
          {
            name: "",
            priority: "",
            comment: ""
          }
        ],
        remarks: ""
      }
    };
  },
  mounted() {
    if(localStorage.name)
    {
      this.model.creatorName = localStorage.name
    }
  },
  methods: {
    add() {
      this.model.path.push({ name: "", priority: "", comment: ""});
    },
        
    handleSubmit(){
            var data = {
            filename : this.model.filename,
            creator_name : this.model.creatorName,
            associations : this.model.path,
            remarks : this.model.remarks
        };
        console.log(data);  
        EmployeeDataService.postDocument(data).then(
            res => {
                // var fileURL = window.URL.createObjectURL(new Blob([res.data]));
                // var fileLink = document.createElement('a');
                // fileLink.href = fileURL;
                // fileLink.download = "qr.png";
                // URL.revokeObjectURL(fileLink.href);
                // console.log(fileLink);
                // fileLink.click();
                console.log(res.data);
                saveAs(blob, 'image_name.jpg');
            }
        ).catch(e => {
            console.log(e)
        });
    },

  }
};
</script>
<style></style>
