<script>
  import Home from "../admin/Home.svelte";
  import dayjs from "dayjs";
  let listAdmin = [];
  let listAdminrule = [];
  let record = "";
  let totalrecord = 0;
  export let table_header_font = "";
  export let table_body_font = "";
  let token = localStorage.getItem("token");
  let akses_page = false;

  async function initapp() {
    const res = await fetch("/api/valid", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        page: "COMPANYADMIN-VIEW",
      }),
    });
    const json = await res.json();
    if (json.status === 400) {
      logout();
    } else if (json.status == 403) {
      alert(json.message);
    } else {
      akses_page = true;
      initAdmin();
    }
  }
  async function initAdmin() {
    const res = await fetch("/api/alladmin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({}),
    });
    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      totalrecord = record.length;
      let recordlistrule = json.listruleadmin;
      let css_status = "";
      let no = 0;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          no = no + 1;
          listAdmin = [
            ...listAdmin,
            {
              admin_no: no,
              admin_id: record[i]["admin_id"],
              admin_idrule: record[i]["admin_idrule"],
              admin_username: record[i]["admin_username"],
              admin_nama: record[i]["admin_nama"],
              admin_rule: record[i]["admin_rule"],
              admin_joindate: record[i]["admin_joindate"],
              admin_lastlogin: record[i]["admin_lastlogin"],
              admin_lastipaddres: record[i]["admin_lastipaddres"],
              admin_status: record[i]["admin_status"],
              admin_status_css: record[i]["admin_status_css"],
            },
          ];
        }
      }
      if (recordlistrule != null) {
        for (let i = 0; i < recordlistrule.length; i++) {
          listAdminrule = [
            ...listAdminrule,
            {
              adminrule_id: recordlistrule[i]["adminrule_id"],
              adminrule_name: recordlistrule[i]["adminrule_name"],
            },
          ];
        }
      }
    } else {
      logout();
    }
  }
  async function logout() {
    localStorage.clear();
    window.location.href = "/";
  }

  const handleRefreshData = (e) => {
    listAdmin = [];
    listAdminrule = [];
    totalrecord = 0;
    setTimeout(function () {
      initAdmin();
    }, 500);
  };
  initapp();
</script>

{#if akses_page == true}
  <Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listAdmin}
    {listAdminrule}
    {totalrecord}
  />
{/if}
