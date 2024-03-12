<script>
    import Button from "../../components/Button.svelte";
    import Input_custom from '../../components/InputCustom.svelte' 
    import Modal from "../../components/Modal.svelte";

    let token = localStorage.getItem("token");
    let path_api = "/";
    let engine_invoice = "";
    let myModal_newentry = "";
    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "COMPANYDASHBOARD-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp()
    var source = new EventSource(path_api+"sse");
    source.onmessage = function(event) {
        let text_dasar = event.data;
        let text_replace1 = text_dasar.replace(`"`,"")
        let text_replace2 = text_replace1.replace(`"`,"")
        let text_finalsplit = text_replace2.split("|");
        console.log(event.data)
        engine_invoice = text_finalsplit[0];
    };

    let flag_btnsave = true;
    let invoice_id_field = ""
    let invoice_date_field = ""
    let invoice_result_field = ""
    let invoice_status_field = ""
    let invoice_statuscss_field = ""
    let invoice_totalbet_field = 0
    let invoice_totalwin_field = 0
    let invoice_winlose_field = 0

    let listallinvoice = []
    let title_allinvoice = "";
    let allinvoice_totalbet = 0;
    let allinvoice_totalwin = 0;
    let allinvoice_winlose = 0;
    let allinvoice_winlose_member = 0;
    let allinvoice_winlose_css = "";
    let allinvoice_winlose_member_css = "";
    let listprediksi = []
    let prediksi_totalmember = 0
    let prediksi_totalbet = 0
    let prediksi_totalwin = 0
    let prediksi_winlose = 0

    let listdetailinvoice = []
    let detail_id = ""
    let detail_date = ""
    let detail_result = ""
    let detail_totalmember = 0
    let detail_totalbet = 0
    let detail_totalwin = 0
    let detail_winlose_agen_css = ""
    let detail_winlose_member_css = ""

    let detail_winloseparent_agen = 0
    let detail_winloseparent_member = 0
    let detail_winloseparent_agen_css = ""
    let detail_winloseparent_member_css = ""
    let detail_tab_win = "active"
    let detail_tab_lose = ""
    let detail_tab_running = ""

    let listPage = [];
    let totalrecord = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let perpage = 0;
    let page = 0;
    let css_loader = "display: none;";
    let msgloader = "";
    const call_allinvoice = () => {
        call_alldatainvoice()
        myModal_newentry = new bootstrap.Modal(document.getElementById("modal_allinvoice"));
        myModal_newentry.show();
    };
    const call_editinvoice = (e) => {
        invoice_id_field = e
        invoice_status_field = ""
        call_invoice(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modal_resultinvoice"));
        myModal_newentry.show();
    };
    const call_detailinvoice_tab = (e) => {
        switch(e){
            case "WIN":
                detail_tab_win = "active"
                detail_tab_lose = ""
                detail_tab_running = ""
                call_invoicedetail(detail_id,"WIN")
                break;
            case "LOSE":
                detail_tab_win = ""
                detail_tab_lose = "active"
                detail_tab_running = ""
                call_invoicedetail(detail_id,"LOSE")
                break;
            case "RUNNING":
                detail_tab_win = ""
                detail_tab_lose = ""
                detail_tab_running = "active"
                call_invoicedetail(detail_id,"RUNNING")
                break;
        }
    };
    const call_detailinvoice = (id,date,result,totalmember,totalbet,totalwin) => {
        detail_id = id
        detail_date = date
        detail_result = result
        detail_totalmember = totalmember
        detail_totalbet = totalbet
        detail_totalwin = totalwin


        detail_winloseparent_agen = 0
        detail_winloseparent_member = 0
        detail_winloseparent_agen_css = ""
        detail_winloseparent_member_css = ""
        detail_winloseparent_agen = parseInt(detail_totalbet) - parseInt(totalwin)
        detail_winloseparent_agen_css = ""
        detail_winlose_member_css = ""
        if(detail_winloseparent_agen>0){
            detail_winloseparent_member = -detail_winloseparent_agen
            detail_winloseparent_agen_css ="color:blue;";
            detail_winloseparent_member_css ="color:red;";
        }else{
            detail_winloseparent_member = Math.abs(detail_winloseparent_agen)
            detail_winloseparent_agen_css ="color:red;";
            detail_winloseparent_member_css ="color:blue;";
        }
        detail_tab_win = "active"
        detail_tab_lose = ""
        detail_tab_running = ""
        call_invoicedetail(id,"WIN")
        myModal_newentry = new bootstrap.Modal(document.getElementById("modal_detailinvoice"));
        myModal_newentry.show();
    };
    const generateNumber = () => {
        let numbergenerate = Math.floor(Math.random() * 100);
        if(numbergenerate < 10){
            numbergenerate = "0"+numbergenerate
        }
        invoice_result_field = numbergenerate
    };
    async function call_invoicedetail(e,s) {
        listdetailinvoice = []
        detail_winlose_agen_css = ""
        detail_winlose_member_css = ""
        const res = await fetch("/api/transaksi2d30sdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                transaksidetail2D30s_invoice: e,
                transaksidetail2D30s_status: s,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let winlose_member = 0;
                let winlose_agen = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    winlose_agen = parseInt(record[i]["transaksi2D30sdetail_bet"])-parseInt(record[i]["transaksi2D30sdetail_win"])
                    if(winlose_agen > 0){
                        winlose_member = -winlose_agen;
                        detail_winlose_agen_css ="color:blue;";
                        detail_winlose_member_css ="color:red;";
                    }else{
                        winlose_member = Math.abs(winlose_agen);
                        detail_winlose_agen_css ="color:red;";
                        detail_winlose_member_css ="color:blue;";
                    }
                    listdetailinvoice = [
                        ...listdetailinvoice,
                        {
                            transaksi2D30sdetail_no: no,
                            transaksi2D30sdetail_id: record[i]["transaksi2D30sdetail_id"],
                            transaksi2D30sdetail_date: record[i]["transaksi2D30sdetail_date"],
                            transaksi2D30sdetail_username: record[i]["transaksi2D30sdetail_username"],
                            transaksi2D30sdetail_ipaddress: record[i]["transaksi2D30sdetail_ipaddress"],
                            transaksi2D30sdetail_device: record[i]["transaksi2D30sdetail_device"],
                            transaksi2D30sdetail_browser: record[i]["transaksi2D30sdetail_browser"],
                            transaksi2D30sdetail_nomor: record[i]["transaksi2D30sdetail_nomor"],
                            transaksi2D30sdetail_tipebet: record[i]["transaksi2D30sdetail_tipebet"],
                            transaksi2D30sdetail_bet: record[i]["transaksi2D30sdetail_bet"],
                            transaksi2D30sdetail_win: record[i]["transaksi2D30sdetail_win"],
                            transaksi2D30sdetail_winlose_agen: winlose_agen,
                            transaksi2D30sdetail_winlose_agen_css: detail_winlose_agen_css,
                            transaksi2D30sdetail_winlose_member: winlose_member,
                            transaksi2D30sdetail_winlose_member_css: detail_winlose_member_css,
                            transaksi2D30sdetail_multiplier: record[i]["transaksi2D30sdetail_multiplier"],
                        },
                    ];
                   
                }
            }
        }
    }
    async function call_invoice(e) {
        const res = await fetch("/api/transaksi2d30s", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                transaksi2D30s_invoice: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    invoice_id_field = record[i]["transaksi2D30s_id"];
                    invoice_date_field = record[i]["transaksi2D30s_date"];
                    invoice_result_field = record[i]["transaksi2D30s_result"];
                    invoice_totalbet_field = record[i]["transaksi2D30s_totalbet"];
                    invoice_totalwin_field = record[i]["transaksi2D30s_totalwin"];
                    invoice_winlose_field = record[i]["transaksi2D30s_winlose"];
                    invoice_status_field = record[i]["transaksi2D30s_status"];
                    invoice_statuscss_field = record[i]["transaksi2D30s_status_css"];
                   
                }
            }
        }
    }
    async function call_alldatainvoice() {
        listallinvoice = []
        allinvoice_totalbet = 0;
        allinvoice_totalwin = 0;
        allinvoice_winlose = 0;
        const res = await fetch("/api/transaksi2d30s", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                transaksi2D30s_page: parseInt(page),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            title_allinvoice = json.periode;
            allinvoice_totalbet = json.totalbet;
            allinvoice_totalwin = json.totalwin;
            allinvoice_winlose = json.winlose_agen;
            allinvoice_winlose_member = json.winlose_member;
            if(allinvoice_winlose > 0){
                allinvoice_winlose_css = "color:blue;"
            }else{
                allinvoice_winlose_css = "color:red;"
            }
            if(allinvoice_winlose_member > 0){
                allinvoice_winlose_member_css = "color:blue;"
            }else{
                allinvoice_winlose_member_css = "color:red;"
            }
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = record.length;
                let no = 0;
                let winlose_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(parseInt(record[i]["transaksi2D30s_winlose"]) > 0){
                        winlose_css = "color:blue;"
                    }else{
                        winlose_css = "color:red;"
                    }
                    listallinvoice = [
                        ...listallinvoice,
                        {
                            transaksi2D30s_no: no,
                            transaksi2D30s_id: record[i]["transaksi2D30s_id"],
                            transaksi2D30s_idcurr: record[i]["transaksi2D30s_idcurr"],
                            transaksi2D30s_date: record[i]["transaksi2D30s_date"],
                            transaksi2D30s_result: record[i]["transaksi2D30s_result"],
                            transaksi2D30s_totalmember: record[i]["transaksi2D30s_totalmember"],
                            transaksi2D30s_totalbet: record[i]["transaksi2D30s_totalbet"],
                            transaksi2D30s_totalwin: record[i]["transaksi2D30s_totalwin"],
                            transaksi2D30s_winlose: record[i]["transaksi2D30s_winlose"],
                            transaksi2D30s_winlose_css: winlose_css,
                            transaksi2D30s_create: record[i]["transaksi2D30s_create"],
                            transaksi2D30s_update: record[i]["transaksi2D30s_update"],
                        },
                    ];
                   
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
            }
        }
    }
    async function call_prediksi() {
        listprediksi = []
        prediksi_totalmember = 0
        prediksi_totalbet = 0
        prediksi_totalwin = 0
        prediksi_winlose = 0
        const res = await fetch("/api/transaksi2d30sprediksi", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                transaksi2D30sprediksi_invoice: invoice_id_field,
                transaksi2D30sprediksi_result: invoice_result_field.toString(),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            prediksi_totalmember = json.totalmember;
            prediksi_totalbet = json.totalbet;
            prediksi_totalwin = json.totalwin;
            prediksi_winlose = json.winlose;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    console.log("asd")
                    listprediksi = [
                        ...listprediksi,
                        {
                            transaksi2D30sprediksi_no: no,
                            transaksi2D30sprediksi_id: record[i]["transaksi2D30sprediksi_id"],
                            transaksi2D30sprediksi_date: record[i]["transaksi2D30sprediksi_date"],
                            transaksi2D30sprediksi_username: record[i]["transaksi2D30sprediksi_username"],
                            transaksi2D30sprediksi_nomor: record[i]["transaksi2D30sprediksi_nomor"],
                            transaksi2D30sprediksi_bet: record[i]["transaksi2D30sprediksi_bet"],
                            transaksi2D30sprediksi_win: record[i]["transaksi2D30sprediksi_win"],
                            transaksi2D30sprediksi_winlose: record[i]["transaksi2D30sprediksi_winlose"],
                        },
                    ];
                   
                }
            }
        }
    }
    async function handleSave_invoice() {
        let flag = true
        let msg = ""
        if(invoice_id_field == ""){
            flag = false
            msg += "The Invoice is required\n"
        }
        if(invoice_result_field == ""){
            flag = false
            msg += "The Result is required\n"
        }
       
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/transaksi2d30ssave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"COMPANY-SAVE",
                    transaksi2D30s_invoice: invoice_id_field,
                    transaksi2D30s_result: invoice_result_field.toString(),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                msgloader = json.message;
                invoice_result_field = ""
                
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    const handleSelectPaging = (event) => {
        page = event.target.value;
        call_alldatainvoice()
    };
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-lg-3">
            <div class="card mt-1">
                <div class="card-header" style="background-color:#f6f6f6 ;padding:0px;">
                    <h5 class="card-title" style="font-size: 30px;padding:0px;">
                        <center>
                            2D 30S
                        </center>
                    </h5>
                </div>
                <div class="card-body overflow-auto" style="padding: 5px;margin:0px;">
                    <center>
                        <span style="font-bold;font-size:20px;">INVOICE</span><br />
                        <span style="font-bold;font-size:30px;">{engine_invoice}</span>
                    </center>
                    <div class="table-responsive">
                        
                    </div>
                </div>
                <div class="card-footer">
                    <center>
                        <Button on:click={() => {
                                call_allinvoice();
                            }} 
                            button_title="<i class='bi bi-file-earmark'></i>&nbsp;Invoice"
                            button_css="btn-primary"/>
                        {#if engine_invoice != ""}
                            <Button on:click={() => {
                                call_editinvoice(engine_invoice);
                            }} 
                            button_title="<i class='bi bi-pencil'></i>&nbsp;Edit"
                            button_css="btn-info"/>
                        {/if}
                        <Button on:click={() => {
                                call_editinvoice(engine_invoice);
                            }} 
                            button_title="<i class='bi bi-gear'></i>&nbsp;Setting"
                            button_css="btn-info"/>
                    </center>
                </div>
            </div>
        </div>
    </div>
</div>  

<Modal
	modal_id="modal_resultinvoice"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{invoice_id_field}"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-5">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Invoice</label>
                    <Input_custom
                        bind:value={invoice_id_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        disabled=disabled
                        input_placeholder="INVOICE"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Result</label>
                    <div class="input-group mb-3">
                        <Input_custom
                            bind:value={invoice_result_field}
                            input_tipe="text_result"
                            input_required="required"
                            input_maxlength="2"
                            input_placeholder="Result"/>
                        <span class="input-group-text">
                            
                            <Button on:click={() => {
                                generateNumber();
                            }} 
                            button_function="SAVE"
                            button_title="<i class='bi bi-gear'></i>&nbsp;&nbsp;Generate"
                            button_css="btn-info"/>
                            &nbsp;
                            <Button on:click={() => {
                                call_prediksi();
                            }} 
                            button_function="SAVE"
                            button_title="<i class='bi bi-gear'></i>&nbsp;&nbsp;Check"
                            button_css="btn-info"/>
                        </span>
                        
                    </div>
                    
                </div>
                <div class="mb-3">
                    <div class="d-grid gap-1">
                        {#if invoice_status_field == "OPEN"}
                            {#if flag_btnsave}
                            <Button on:click={() => {
                                    handleSave_invoice();
                                }} 
                                button_function="SAVE"
                                button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
                                button_css="btn-warning"/>
                            {/if}
                        {/if}
                    </div>
                </div>
                
            </div>
            <div class="col-md-7">
                <table class="table">
                    <tr>
                        <td>Total Member</td>
                        <td>:</td>
                        <td style="color:blue;">{new Intl.NumberFormat().format(prediksi_totalmember)}</td>
                    </tr>
                    <tr>
                        <td>Total Bet</td>
                        <td>:</td>
                        <td style="color:blue;">{new Intl.NumberFormat().format(prediksi_totalbet)}</td>
                    </tr>
                    <tr>
                        <td>Total Win</td>
                        <td>:</td>
                        <td style="color:blue;">{new Intl.NumberFormat().format(prediksi_totalwin)}</td>
                    </tr>
                    <tr>
                        <td>Winlose</td>
                        <td>:</td>
                        <td style="color:blue;">{new Intl.NumberFormat().format(prediksi_winlose)}</td>
                    </tr>
                </table>
                <table class="table table-sm">
                    <thead>
                        <tr>
                            <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:12px;">USERNAME</th>
                            <th width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:12px;">NOMOR</th>
                            <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">BET</th>
                            <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WIN</th>
                            <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WINLOSE</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each listprediksi as rec}
                        <tr>
                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: 12px;">{rec.transaksi2D30sprediksi_username}</td>
                            <td NOWRAP style="text-align: center;vertical-align: top;font-size: 12px;">{rec.transaksi2D30sprediksi_nomor}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30sprediksi_bet)}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30sprediksi_win)}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;">{new Intl.NumberFormat().format(rec.transaksi2D30sprediksi_winlose)}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
        </div>
	</slot:template>
</Modal>

<Modal
	modal_id="modal_allinvoice"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="PERIODE : {title_allinvoice}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <div class="float-end">
            <select
                on:change={handleSelectPaging}
                style="text-align: center;"
                class="form-control">
                {#each listPage as rec}
                    <option value={rec.page_value}>{rec.page_display}</option>
                {/each}
            </select>
        </div>
    </slot:template>
	<slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:12px;">INVOICE</th>
                    <th width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:12px;">DATE</th>
                    <th width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:12px;">NOMOR</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">MEMBER</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">BET</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WIN</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WINLOSE</th>
                </tr>
            </thead>
            <tbody>
                {#each listallinvoice as rec}
                <tr>
                    <td on:click={() => {
                        //id,result,totalmember,totalbet,totalwin
                        call_detailinvoice(rec.transaksi2D30s_id,rec.transaksi2D30s_date,rec.transaksi2D30s_result,
                        rec.transaksi2D30s_totalmember,rec.transaksi2D30s_totalbet,rec.transaksi2D30s_totalwin);
                    }} NOWRAP 
                        style="text-align: left;vertical-align: top;font-size: 12px;text-decoration: underline;cursor: pointer;">{rec.transaksi2D30s_id}</td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: 12px;">{rec.transaksi2D30s_date}</td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: 12px;">{rec.transaksi2D30s_result}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30s_totalmember)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30s_totalbet)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{new Intl.NumberFormat().format(rec.transaksi2D30s_totalwin)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;{rec.transaksi2D30s_winlose_css}">{new Intl.NumberFormat().format(rec.transaksi2D30s_winlose)}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
    <slot:template slot="footer">
        <table class="table">
            <tr>
                <td style="font-size: 12px;text-align: right;">TOTAL BET</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;color:blue;text-align: right;">{new Intl.NumberFormat().format(allinvoice_totalbet)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: right;">TOTAL WIN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;color:red;text-align: right;">{new Intl.NumberFormat().format(allinvoice_totalwin)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: right;">WINLOSE AGEN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;{allinvoice_winlose_css}text-align: right;">{new Intl.NumberFormat().format(allinvoice_winlose)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: right;">WINLOSE MEMBER</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;{allinvoice_winlose_member_css}text-align: right;">{new Intl.NumberFormat().format(allinvoice_winlose_member)}</td>
            </tr>
        </table>
	</slot:template>
</Modal>

<Modal
	modal_id="modal_detailinvoice"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{detail_id}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table">
            <tr>
                <td style="font-size: 12px;">Tanggal</td>
                <td style="font-size: 12px;">:</td>
                <td style="font-size: 12px;text-align: center;">{detail_date}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;">Result</td>
                <td style="font-size: 12px;">:</td>
                <td style="font-size: 12px;text-align: center;">{detail_result}</td>
            </tr>
           
        </table>
        <ul class="nav nav-pills">
            <li on:click={() => {
                call_detailinvoice_tab("WIN");
            }} class="nav-item" style="cursor: pointer;">
              <span class="nav-link {detail_tab_win}">Win</span>
            </li>
            <li on:click={() => {
                call_detailinvoice_tab("LOSE");
            }} class="nav-item" style="cursor: pointer;">
              <span class="nav-link {detail_tab_lose}">Lose</span>
            </li>
            <li on:click={() => {
                call_detailinvoice_tab("RUNNING");
            }} class="nav-item" style="cursor: pointer;">
              <span class="nav-link {detail_tab_running}" >Running</span>
            </li>
        </ul>
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:12px;">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:12px;">USERNAME</th>
                    <th width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:12px;">NOMOR</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">BET</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WIN</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WINLOSE MEMBER</th>
                    <th width="20%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:12px;">WINLOSE AGEN</th>
                </tr>
            </thead>
            <tbody>
                {#each listdetailinvoice as rec}
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: 12px;">{rec.transaksi2D30sdetail_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: 12px;">{rec.transaksi2D30sdetail_username}</td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: 12px;">{rec.transaksi2D30sdetail_nomor}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30sdetail_bet)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{new Intl.NumberFormat().format(rec.transaksi2D30sdetail_win)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;{rec.transaksi2D30sdetail_winlose_member_css}">{new Intl.NumberFormat().format(rec.transaksi2D30sdetail_winlose_member)}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: 12px;{rec.transaksi2D30sdetail_winlose_agen_css}">{new Intl.NumberFormat().format(rec.transaksi2D30sdetail_winlose_agen)}</td>
                </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
    <slot:template slot="footer">
        <table class="table">
            <tr>
                <td style="font-size: 12px;">Total Member</td>
                <td style="font-size: 12px;">:</td>
                <td style="color:blue;font-size: 12px;text-align: right;">{new Intl.NumberFormat().format(detail_totalmember)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;">Total Bet</td>
                <td style="font-size: 12px;">:</td>
                <td style="color:blue;font-size: 12px;text-align: right;">{new Intl.NumberFormat().format(detail_totalbet)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;">Total Win</td>
                <td style="font-size: 12px;">:</td>
                <td style="color:blue;font-size: 12px;text-align: right;">{new Intl.NumberFormat().format(detail_totalwin)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;">Winlose Agen</td>
                <td style="font-size: 12px;">:</td>
                <td style="{detail_winloseparent_agen_css}font-size: 12px;text-align: right;">{new Intl.NumberFormat().format(detail_winloseparent_agen)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;">Winlose Member</td>
                <td style="font-size: 12px;">:</td>
                <td style="{detail_winloseparent_member_css}font-size: 12px;text-align: right;">{new Intl.NumberFormat().format(detail_winloseparent_member)}</td>
            </tr>
        </table>
	</slot:template>
</Modal>