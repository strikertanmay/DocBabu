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
            <h1 class="display-2 text-white">Update Track</h1>
            <p class="text-white mt-0 mb-5">Update track for the processing of a file</p>
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
                        label="File ID"
                        placeholder="enter file id"
                        input-classes="form-control-alternative"
                        v-model="model.id"
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
                      <base-input
                        alternative
                        label="Received time"
                        placeholder="time for creation of track"
                        input-classes="form-control-alternative"
                        v-model="model.receiveTime"
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
export default {
  name: "user-profile",
  data() {
    return {
      model: {
        id: "",
        receivedName: "",
        receiveTime: "",
        status: "",
        send: ""
      }
    };
  }
};
</script>
<style></style>
