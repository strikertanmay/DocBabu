<template>
  <div class="card shadow"
       :class="type === 'dark' ? 'bg-default': ''">
    <div class="card-header border-0"
         :class="type === 'dark' ? 'bg-transparent': ''">
      <div class="row align-items-center">
        <div class="col">
          <h3 class="mb-0" :class="type === 'dark' ? 'text-white': ''">
            {{title}}
          </h3>
        </div>
        <div class="col text-right">
          <!-- <base-button type="primary" size="sm">See all</base-button> -->
        </div>
      </div>
    </div>

    <div class="table-responsive">
      <base-table class="table align-items-center table-flush"
                  :class="type === 'dark' ? 'table-dark': ''"
                  :thead-classes="type === 'dark' ? 'thead-dark': 'thead-light'"
                  tbody-classes="list"
                  :data="tableData">
        <template slot="columns">
          <th>File Identifier</th>
          <th>Recieved From</th>
          <th>Status</th>
          <th>Send To</th>
          <th>Last Modified</th>
          <th>Action</th>
          <th></th>
        </template>

        <template slot-scope="{row}">
          <th scope="row">
            <div class="media align-items-center">
              <a href="#" class="avatar rounded-circle mr-3">
                <img alt="Image placeholder" :src="row.img">
              </a>
              <div class="media-body">
                <span class="name mb-0 text-sm">{{row.title}}</span>
              </div>
            </div>
          </th>
          <td class="budget">
            {{row.receive}}
          </td>
<!--         
          <td>
            <div class="avatar-group">
              <a href="#" class="avatar avatar-sm rounded-circle" data-toggle="tooltip" data-original-title="Ryan Tompson">
                <img alt="Image placeholder" src="img/theme/team-1-800x800.jpg">
              </a>
              <a href="#" class="avatar avatar-sm rounded-circle" data-toggle="tooltip" data-original-title="Romina Hadid">
                <img alt="Image placeholder" src="img/theme/team-2-800x800.jpg">
              </a>
              <a href="#" class="avatar avatar-sm rounded-circle" data-toggle="tooltip" data-original-title="Alexander Smith">
                <img alt="Image placeholder" src="img/theme/team-3-800x800.jpg">
              </a>
              <a href="#" class="avatar avatar-sm rounded-circle" data-toggle="tooltip" data-original-title="Jessica Doe">
                <img alt="Image placeholder" src="img/theme/team-4-800x800.jpg">
              </a>
            </div>
          </td> -->
            <td>
            <badge class="badge-dot mr-4" :type="row.statusType">
              <i :class="`bg-${row.statusType}`"></i>
              <span class="status">{{row.status}}</span>
            </badge>
          </td>
          <td>
            <div class="d-flex align-items-center">
              <span class="completion mr-2">{{row.send}}</span>
              <!-- <div>
                <base-progress :type="row.statusType"
                               :show-percentage="false"
                               class="pt-0"
                               :value="row.completion"/>
              </div> -->
            </div>
          </td>
          <td class="budget">
            {{row.lastmod}}
          </td>
          <td>
             <base-button type="primary" size="sm">
              <router-link class="nav-link nav-link-icon" to="/update" style="color:white">
                Update
              </router-link>
            </base-button>
          </td>

        

        </template>

      </base-table>
    </div>

    <div class="card-footer d-flex justify-content-end"
         :class="type === 'dark' ? 'bg-transparent': ''">
      <base-pagination total="30"></base-pagination>
    </div>

  </div>
</template>
<script>
  export default {
    name: 'projects-table',
    props: {
      type: {
        type: String
      },
      title: String
    },
    data() {
      return {
        tableData: [
          {
            img: 'img/theme/bootstrap.jpg',
            title: 'Academic_Calander 20',
            receive: 'Gymkhana', 
            status: 'on schedule',
            statusType: 'info',
            send: 'AAA',
            lastmod: '25 January 2020'
          },
          {
            img: 'img/theme/bootstrap.jpg',
            title: 'Permission for CC3',
            receive: 'GeekHaven',
            status: 'delayed',
            statusType: 'danger',
            send: 'Admin',
            lastmod: '3 February 2020'
          },
          {
            img: 'img/theme/bootstrap.jpg',
            title: 'Annual Budget',
            receive: 'Finance Team',
            status: 'completed',
            statusType: 'success',
            send: 'Director',
            lastmod: '6 January 2020'
          },
          {
            img: 'img/theme/bootstrap.jpg',
            title: 'Google ML Workshop',
            receive: 'GeekHaven',
            status: 'pending',
            statusType: 'warning',
            send: 'GeekHaven',
            lastmod: '17 January 2020'
          },
          {
            img: 'img/theme/bootstrap.jpg',
            title: 'Semester Time Table',
            receive: 'Gymkhana',
            status: 'completed',
            statusType: 'success',
            send: 'Students',
            lastmod: '12 December 2019'
          }
        ]
      }
    }
  }
</script>
<style>
</style>
