<template>
    <div>
        <base-header type="gradient-success" class="pb-6 pb-8 pt-5 pt-md-8">
            <!-- Card stats -->
                        <div class="row">
                <div class="col-xl-3 col-lg-6">
                    <stats-card title="Total Files Pending"
                                type="gradient-red"
                                sub-title="32"
                                icon="ni ni-active-40"
                                class="mb-4 mb-xl-0"
                    >

                        <template slot="footer">
                            <!-- <span class="text-success mr-2"><i class="fa fa-arrow-up"></i> </span> -->
                            <span class="text-nowrap">Since last month</span>
                        </template>
                    </stats-card>
                </div>
                <div class="col-xl-3 col-lg-6">
                    <stats-card title="Last File Signed"
                                type="gradient-orange"
                                sub-title="Aparoksha_Budget'19"
                                icon="ni ni-chart-pie-35"
                                class="mb-4 mb-xl-0"
                    >

                        <template slot="footer">
                            <!-- <span class="text-success mr-2"><i class="fa fa-arrow-up"></i> </span> -->
                            <span class="text-nowrap">2 days ago</span>
                        </template>
                    </stats-card>
                </div>
                <div class="col-xl-3 col-lg-6">
                    <stats-card title="Latest File Pending"
                                type="gradient-green"
                                sub-title="Academic_Calender'20"
                                icon="ni ni-money-coins"
                                class="mb-4 mb-xl-0"
                    >

                        <template slot="footer">
                            <!-- <span class="text-danger mr-2"><i class="fa fa-arrow-down"></i></span> -->
                            <span class="text-nowrap">Since 1 day</span>
                        </template>
                    </stats-card>

                </div>
                <div class="col-xl-3 col-lg-6">
                    <stats-card title="Files COmpleted this Week"
                                type="gradient-info"
                                sub-title="12"
                                icon="ni ni-chart-bar-32"
                                class="mb-4 mb-xl-0"
                    >

                        <template slot="footer">
                            <!-- <span class="text-success mr-2"><i class="fa fa-arrow-up"></i> 54.8%</span> -->
                            <span class="text-nowrap">Since last week</span>
                        </template>
                    </stats-card>
                </div>
            </div>
        </base-header>

        <div class="container-fluid mt--7">
            <!-- <div class="row">
                <div class="col">
                    <projects-table title="Light Table"></projects-table>
                </div>
            </div> -->
            <div class="row mt-5">
                <div class="col">
                    <projects-table type="dark" title="Dark Table"></projects-table>
                </div>
            </div>
        </div>

    </div>
</template>
<script>
  import ProjectsTable from './Tables/ProjectsTable'
  export default {
    name: 'tables',
    components: {
      ProjectsTable
    }
  };
</script>
<style></style>
