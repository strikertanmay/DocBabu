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
            <h1 class="display-2 text-white">Update A Track</h1>
            <p class="text-white mt-0 mb-5" style="font-size:20px">Update the track of a file in your hands. Begin with scanning the QR and stating the file information along with the sender's and recipient's details.
              Mark the status as complete and specify the next authority to process the document further. The time for the latest modification and other details will be displayed to associated employees.
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
                  <h3 class="mb-0">Update Track</h3>
                </div>
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
                        label="File Name"
                        placeholder="enter file name"
                        input-classes="form-control-alternative"
                        v-model="model.filename"
                      />
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-lg-6">
                      <base-input
                        alternative
                        label="Received From"
                        placeholder="employee name"
                        input-classes="form-control-alternative"
                        v-model="model.receivedName"
                      />
                    </div>
                    <div class="col-lg-6">
                      <input
                        type="file"
                        accept="image/*" 
                        @change="uploadImage" 
                        id="file-input"
                        alternative
                        label="QR Code"
                        placeholder="upload qr code for file"
                        input-classes="form-control-alternative"
                      />
                    </div>
                  </div>
                </div>
                <hr class="my-4" />
                <!-- Address -->
                <h6 class="heading-small text-muted mb-4">Track Information</h6>
                <div class="pl-lg-4">
                  <div class="row">
                    <div class="col-lg-6">
                      <base-input
                        alternative
                        label="Completion Status"
                        placeholder="status"
                        input-classes="form-control-alternative"
                        v-model="model.status"
                      />
                    </div>
                    <div class="col-lg-6">
                      <base-input
                        alternative
                        label="Sent To"
                        placeholder="employee name"
                        input-classes="form-control-alternative"
                        v-model="model.send"
                      />
                    </div>
                  </div>
                </div>
                <hr class="my-4" />
                <!-- Description -->
                <h6 class="heading-small text-muted mb-4">Additional Information</h6>
                <div class="pl-lg-4">
                  <div class="form-group">
                    <base-input alternative label="Remarks">
                      <textarea
                        rows="4"
                        class="form-control form-control-alternative"
                        placeholder="A few words about you ..."
                      >A beautiful Dashboard for Bootstrap 4. It is Free and Open Source.</textarea>
                    </base-input>
                  </div>
                </div>
                <div class="pl-lg-4" >
                  <base-button type="primary" class="my-4" @click="handleSubmit">Submit</base-button>
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
        receivedName: "",
        receiveTime: "",
        status: "",
        send: "",
        qr: null
      }
    };
  },
  methods: {
    uploadImage(e){
                const image = e.target.files[0];
                const reader = new FileReader();
                reader.readAsDataURL(image);
                reader.onload = e =>{
                    this.model.qr = e.target.result;
                    console.log(this.model.qr);
                };
            },
    handleSubmit() {
            var data = {
            user_id : localStorage.id,
            image : IMAGE,
        };
        console.log(data);  
        EmployeeDataService.postUpdate(data).then(
            res => {
                console.log(res.data);
            }
        ).catch(e => {
            console.log(e)
        });
    }
  }
};
</script>
<style></style>
