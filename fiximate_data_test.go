package quickfixdev

const html0 = `<html xmlns="http://www.w3.org/1999/xhtml" lang="en">

<head>
    <link rel="stylesheet" href="../../fixstyle.css" type="text/css"></link>
    <title>FIX.5.0SP2_EP252 - Fields sorted by Tag Number</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"></meta>
</head>

<body>
    <h2>FIX.5.0SP2_EP252 - Fields sorted by Tag Number</h2>
    <table class="ListTable">
        <tr>
            <th class="FldTagHdr">Tag</th>
            <th class="FldNameHdr">Field Name</th>
            <th class="FldNameHdr">XML Name</th>
            <th class="FldDatHdr">Data Type</th>
            <th class="FldDatHdr">Union Datatype</th>
            <th class="FldDesHdr">Description</th>
            <th class="FldDesHdr">Added</th>
            <th class="FldDesHdr">Depr.</th>
            <th class="FldDesHdr">Enums from tag</th>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag1.html" target="tagFrame">1</a></td>
            <td class="FldNameOdd"><a href="tag1.html" target="tagFrame">Account</a></td>
            <td class="FldNameOdd">
                @Acct</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Account mnemonic as agreed between buy and sell sides, e.g. broker and institution or investor/intermediary and fund manager.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag2.html" target="tagFrame">2</a></td>
            <td class="FldNameEven"><a href="tag2.html" target="tagFrame">AdvId</a></td>
            <td class="FldNameEven">
                @AdvId</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Unique identifier of advertisement message.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag3.html" target="tagFrame">3</a></td>
            <td class="FldNameOdd"><a href="tag3.html" target="tagFrame">AdvRefID</a></td>
            <td class="FldNameOdd">
                @AdvRefID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Reference identifier used with CANCEL and REPLACE transaction types.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag4.html" target="tagFrame">4</a></td>
            <td class="FldNameEven"><a href="tag4.html" target="tagFrame">AdvSide</a></td>
            <td class="FldNameEven">
                @AdvSide</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Broker's side of advertised trade</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag5.html" target="tagFrame">5</a></td>
            <td class="FldNameOdd"><a href="tag5.html" target="tagFrame">AdvTransType</a></td>
            <td class="FldNameOdd">
                @AdvTransTyp</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies advertisement message transaction type</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag6.html" target="tagFrame">6</a></td>
            <td class="FldNameEven"><a href="tag6.html" target="tagFrame">AvgPx</a></td>
            <td class="FldNameEven">
                @AvgPx</td>
            <td class="FldDatEven">Price</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Calculated average price of all fills on this order.</span></p><span class="">
         </span>
                <p class="Even"><span class="">For Fixed Income trades AvgPx is always expressed as percent-of-par, regardless of the PriceType (423) of LastPx (31). I.e., AvgPx will contain an average of percent-of-par values (see LastParPx (669)) for issues traded in Yield, Spread or Discount.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag7.html" target="tagFrame">7</a></td>
            <td class="FldNameOdd"><a href="tag7.html" target="tagFrame">BeginSeqNo</a></td>
            <td class="FldNameOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatOdd">SeqNum</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Message sequence number of first message in range to be resent</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag8.html" target="tagFrame">8</a></td>
            <td class="FldNameEven"><a href="tag8.html" target="tagFrame">BeginString</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Identifies beginning of new message and protocol version. ALWAYS FIRST FIELD IN MESSAGE. (Always unencrypted)</span></p><span class="">
         </span>
                <p class="Even"><span class="">Valid values:</span></p><span class="">
         </span>
                <p class="Even"><span class="">FIXT.1.1</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag9.html" target="tagFrame">9</a></td>
            <td class="FldNameOdd"><a href="tag9.html" target="tagFrame">BodyLength</a></td>
            <td class="FldNameOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatOdd">Length</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Message length, in bytes, forward to the CheckSum field. ALWAYS SECOND FIELD IN MESSAGE. (Always unencrypted)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag10.html" target="tagFrame">10</a></td>
            <td class="FldNameEven"><a href="tag10.html" target="tagFrame">CheckSum</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Three byte, simple checksum (see Volume 2: "Checksum Calculation" for description). ALWAYS LAST FIELD IN MESSAGE; i.e. serves, with the trailing &lt;SOH&gt;, as the end-of-message delimiter. Always defined as three characters. (Always unencrypted)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag11.html" target="tagFrame">11</a></td>
            <td class="FldNameOdd"><a href="tag11.html" target="tagFrame">ClOrdID</a></td>
            <td class="FldNameOdd">
                @ClOrdID
                <br> @ID in SingleGeneralOrderHandling</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Unique identifier for Order as assigned by the buy-side (institution, broker, intermediary etc.) (identified by SenderCompID (49) or OnBehalfOfCompID (5) as appropriate). Uniqueness must be guaranteed within a single trading day. Firms, particularly those which electronically submit multi-day orders, trade globally or throughout market close periods, should ensure uniqueness across days, for example by embedding a date within the ClOrdID field.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag12.html" target="tagFrame">12</a></td>
            <td class="FldNameEven"><a href="tag12.html" target="tagFrame">Commission</a></td>
            <td class="FldNameEven">
                @Comm</td>
            <td class="FldDatEven">Amt</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Commission. Note if CommType (13) is percentage, Commission of 5% should be represented as .05.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag13.html" target="tagFrame">13</a></td>
            <td class="FldNameOdd"><a href="tag13.html" target="tagFrame">CommType</a></td>
            <td class="FldNameOdd">
                @CommTyp</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Specifies the basis or unit used to calculate the total commission based on the rate.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP2 &nbsp;EP204
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag14.html" target="tagFrame">14</a></td>
            <td class="FldNameEven"><a href="tag14.html" target="tagFrame">CumQty</a></td>
            <td class="FldNameEven">
                @CumQty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Total quantity (e.g. number of shares) filled.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag15.html" target="tagFrame">15</a></td>
            <td class="FldNameOdd"><a href="tag15.html" target="tagFrame">Currency</a></td>
            <td class="FldNameOdd">
                @Ccy</td>
            <td class="FldDatOdd">Currency</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies currency used for price. Absence of this field is interpreted as the default for the security. It is recommended that systems provide the currency value whenever possible. See "Appendix 6-A: Valid Currency Codes" for information on obtaining valid values.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag16.html" target="tagFrame">16</a></td>
            <td class="FldNameEven"><a href="tag16.html" target="tagFrame">EndSeqNo</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">SeqNum</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Message sequence number of last message in range to be resent. If request is for a single message BeginSeqNo (7) = EndSeqNo. If request is for all messages subsequent to a particular message, EndSeqNo = "0" (representing infinity).</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag17.html" target="tagFrame">17</a></td>
            <td class="FldNameOdd"><a href="tag17.html" target="tagFrame">ExecID</a></td>
            <td class="FldNameOdd">
                @ExecID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Unique identifier of execution message as assigned by sell-side (broker, exchange, ECN) (will be 0 (zero) for ExecType (150)=I (Order Status)).</span></p><span class="">
         </span>
                <p class="Odd"><span class="">Uniqueness must be guaranteed within a single trading day or the life of a multi-day order. Firms which accept multi-day orders should consider embedding a date within the ExecID field to assure uniqueness across days.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.1 this field was of type int).</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP1 &nbsp;EP95
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag18.html" target="tagFrame">18</a></td>
            <td class="FldNameEven"><a href="tag18.html" target="tagFrame">ExecInst</a></td>
            <td class="FldNameEven">
                @ExecInst</td>
            <td class="FldDatEven">MultipleCharValue</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Instructions for order handling on exchange trading floor. If more than one instruction is applicable to an order, this field can contain multiple instructions separated by space. *** SOME VALUES HAVE BEEN REPLACED - See "Replaced Features and Supported Approach" *** (see Volume : "Glossary" for value definitions)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag19.html" target="tagFrame">19</a></td>
            <td class="FldNameOdd"><a href="tag19.html" target="tagFrame">ExecRefID</a></td>
            <td class="FldNameOdd">
                @ExecRefID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Reference identifier used with Trade, Trade Cancel and Trade Correct execution types.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag21.html" target="tagFrame">21</a></td>
            <td class="FldNameEven"><a href="tag21.html" target="tagFrame">HandlInst</a></td>
            <td class="FldNameEven">
                @HandlInst</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Instructions for order handling on Broker trading floor</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag22.html" target="tagFrame">22</a></td>
            <td class="FldNameOdd"><a href="tag22.html" target="tagFrame">SecurityIDSource</a></td>
            <td class="FldNameOdd">
                @Src</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd">Reserved100Plus</td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies class or source of the SecurityID(48) value.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP2 &nbsp;EP161
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag23.html" target="tagFrame">23</a></td>
            <td class="FldNameEven"><a href="tag23.html" target="tagFrame">IOIID</a></td>
            <td class="FldNameEven">
                @IOIID
                <br> @ID in Indication</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Unique identifier of IOI message.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag25.html" target="tagFrame">25</a></td>
            <td class="FldNameOdd"><a href="tag25.html" target="tagFrame">IOIQltyInd</a></td>
            <td class="FldNameOdd">
                @QltyInd</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Relative quality of indication</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag26.html" target="tagFrame">26</a></td>
            <td class="FldNameEven"><a href="tag26.html" target="tagFrame">IOIRefID</a></td>
            <td class="FldNameEven">
                @RefID</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Reference identifier used with CANCEL and REPLACE, transaction types.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag27.html" target="tagFrame">27</a></td>
            <td class="FldNameOdd"><a href="tag27.html" target="tagFrame">IOIQty</a></td>
            <td class="FldNameOdd">
                @Qty</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd">Qty</td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Quantity (e.g. number of shares) in numeric form or relative size.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag28.html" target="tagFrame">28</a></td>
            <td class="FldNameEven"><a href="tag28.html" target="tagFrame">IOITransType</a></td>
            <td class="FldNameEven">
                @TransTyp</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Identifies IOI message transaction type</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag29.html" target="tagFrame">29</a></td>
            <td class="FldNameOdd"><a href="tag29.html" target="tagFrame">LastCapacity</a></td>
            <td class="FldNameOdd">
                @LastCpcty</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Broker capacity in order execution</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag30.html" target="tagFrame">30</a></td>
            <td class="FldNameEven"><a href="tag30.html" target="tagFrame">LastMkt</a></td>
            <td class="FldNameEven">
                @LastMkt</td>
            <td class="FldDatEven">Exchange</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Market of execution for last fill, or an indication of the market where an order was routed</span></p><span class="">
         </span>
                <p class="Even"><span class="">Valid values:</span></p><span class="">
         </span>
                <p class="Even"><span class="">See "Appendix 6-C"</span></p><span class="">
      </span>
                <br><i><span class="">
         </span><p class="Even"><span class="">In the context of ESMA RTS 1 Annex I, Table 3, Field 6 "Venue of Execution" it is required that the "venue where the transaction was executed" be identified using ISO 10383 (MIC). Additionally, ESMA requires the use of "MIC code 'XOFF' for financial instruments admitted to trading or traded on a trading venue, where the transaction on that financial instrument is not executed on a trading venue, systematic internaliser or organized trading platform outside of the Union. Use 'SINT' for financial instruments admitted to trading or traded on a trading venue, where the transaction is executed on a systematic internaliser."</span></p><span class="">
      </span></i>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP2 &nbsp;EP228
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag31.html" target="tagFrame">31</a></td>
            <td class="FldNameOdd"><a href="tag31.html" target="tagFrame">LastPx</a></td>
            <td class="FldNameOdd">
                @LastPx</td>
            <td class="FldDatOdd">Price</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Price of this (last) fill.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag32.html" target="tagFrame">32</a></td>
            <td class="FldNameEven"><a href="tag32.html" target="tagFrame">LastQty</a></td>
            <td class="FldNameEven">
                @LastQty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Quantity (e.g. shares) bought/sold on this (last) fill.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag33.html" target="tagFrame">33</a></td>
            <td class="FldNameOdd"><a href="tag33.html" target="tagFrame">NoLinesOfText</a></td>
            <td class="FldNameOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatOdd">NumInGroup</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies number of lines of text body</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag34.html" target="tagFrame">34</a></td>
            <td class="FldNameEven"><a href="tag34.html" target="tagFrame">MsgSeqNum</a></td>
            <td class="FldNameEven">
                @SeqNum</td>
            <td class="FldDatEven">SeqNum</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Integer message sequence number.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag35.html" target="tagFrame">35</a></td>
            <td class="FldNameOdd"><a href="tag35.html" target="tagFrame">MsgType</a></td>
            <td class="FldNameOdd">
                @MsgTyp</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Defines message type ALWAYS THIRD FIELD IN MESSAGE. (Always unencrypted)</span></p><span class="">
         </span>
                <p class="Odd"><span class="">Note: A "U" as the first character in the MsgType field (i.e. U, U2, etc) indicates that the message format is privately defined between the sender and receiver.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">*** Note the use of lower case letters ***</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag36.html" target="tagFrame">36</a></td>
            <td class="FldNameEven"><a href="tag36.html" target="tagFrame">NewSeqNo</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">SeqNum</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">New sequence number</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag37.html" target="tagFrame">37</a></td>
            <td class="FldNameOdd"><a href="tag37.html" target="tagFrame">OrderID</a></td>
            <td class="FldNameOdd">
                @OrdID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Unique identifier for Order as assigned by sell-side (broker, exchange, ECN). Uniqueness must be guaranteed within a single trading day. Firms which accept multi-day orders should consider embedding a date within the OrderID field to assure uniqueness across days.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag38.html" target="tagFrame">38</a></td>
            <td class="FldNameEven"><a href="tag38.html" target="tagFrame">OrderQty</a></td>
            <td class="FldNameEven">
                @Qty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Quantity ordered. This represents the number of shares for equities or par, face or nominal value for FI instruments.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag39.html" target="tagFrame">39</a></td>
            <td class="FldNameOdd"><a href="tag39.html" target="tagFrame">OrdStatus</a></td>
            <td class="FldNameOdd">
                @OrdStat
                <br> @Stat in SingleGeneralOrderHandling</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies current status of order. *** SOME VALUES HAVE BEEN REPLACED - See "Replaced Features and Supported Approach" *** (see Volume : "Glossary" for value definitions)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag40.html" target="tagFrame">40</a></td>
            <td class="FldNameEven"><a href="tag40.html" target="tagFrame">OrdType</a></td>
            <td class="FldNameEven">
                @OrdTyp
                <br> @Typ in SingleGeneralOrderHandling</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Order type. *** SOME VALUES ARE NO LONGER USED - See "Deprecated (Phased-out) Features and Supported Approach" *** (see Volume : "Glossary" for value definitions)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag41.html" target="tagFrame">41</a></td>
            <td class="FldNameOdd"><a href="tag41.html" target="tagFrame">OrigClOrdID</a></td>
            <td class="FldNameOdd">
                @OrigClOrdID
                <br> @OrigID in SingleGeneralOrderHandling</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">ClOrdID (11) of the previous order (NOT the initial order of the day) as assigned by the institution, used to identify the previous order in cancel and cancel/replace requests.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag42.html" target="tagFrame">42</a></td>
            <td class="FldNameEven"><a href="tag42.html" target="tagFrame">OrigTime</a></td>
            <td class="FldNameEven">
                @OrigTm</td>
            <td class="FldDatEven">UTCTimestamp</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Time of message origination (always expressed in UTC (Universal Time Coordinated, also known as "GMT"))</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag43.html" target="tagFrame">43</a></td>
            <td class="FldNameOdd"><a href="tag43.html" target="tagFrame">PossDupFlag</a></td>
            <td class="FldNameOdd">
                @PosDup</td>
            <td class="FldDatOdd">Boolean</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Indicates possible retransmission of message with this sequence number</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag44.html" target="tagFrame">44</a></td>
            <td class="FldNameEven"><a href="tag44.html" target="tagFrame">Price</a></td>
            <td class="FldNameEven">
                @Px</td>
            <td class="FldDatEven">Price</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Price per unit of quantity (e.g. per share)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag45.html" target="tagFrame">45</a></td>
            <td class="FldNameOdd"><a href="tag45.html" target="tagFrame">RefSeqNum</a></td>
            <td class="FldNameOdd">
                @RefSeqNum</td>
            <td class="FldDatOdd">SeqNum</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Reference message sequence number</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag48.html" target="tagFrame">48</a></td>
            <td class="FldNameEven"><a href="tag48.html" target="tagFrame">SecurityID</a></td>
            <td class="FldNameEven">
                @ID</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Security identifier value of SecurityIDSource (22) type (e.g. CUSIP, SEDOL, ISIN, etc). Requires SecurityIDSource.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag49.html" target="tagFrame">49</a></td>
            <td class="FldNameOdd"><a href="tag49.html" target="tagFrame">SenderCompID</a></td>
            <td class="FldNameOdd">
                @SID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Assigned value used to identify firm sending message.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag50.html" target="tagFrame">50</a></td>
            <td class="FldNameEven"><a href="tag50.html" target="tagFrame">SenderSubID</a></td>
            <td class="FldNameEven">
                @SSub</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Assigned value used to identify specific message originator (desk, trader, etc.)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag52.html" target="tagFrame">52</a></td>
            <td class="FldNameOdd"><a href="tag52.html" target="tagFrame">SendingTime</a></td>
            <td class="FldNameOdd">
                @Snt</td>
            <td class="FldDatOdd">UTCTimestamp</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Time of message transmission (always expressed in UTC (Universal Time Coordinated, also known as "GMT")</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag53.html" target="tagFrame">53</a></td>
            <td class="FldNameEven"><a href="tag53.html" target="tagFrame">Quantity</a></td>
            <td class="FldNameEven">
                @Qty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Overall/total quantity (e.g. number of shares)</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag54.html" target="tagFrame">54</a></td>
            <td class="FldNameOdd"><a href="tag54.html" target="tagFrame">Side</a></td>
            <td class="FldNameOdd">
                @Side</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Side of order (see Volume : "Glossary" for value definitions)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag55.html" target="tagFrame">55</a></td>
            <td class="FldNameEven"><a href="tag55.html" target="tagFrame">Symbol</a></td>
            <td class="FldNameEven">
                @Sym</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Ticker symbol. Common, "human understood" representation of the security. SecurityID (48) value can be specified if no symbol exists (e.g. non-exchange traded Collective Investment Vehicles)</span></p><span class="">
         </span>
                <p class="Even"><span class="">Use "[N/A]" for products which do not have a symbol.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag56.html" target="tagFrame">56</a></td>
            <td class="FldNameOdd"><a href="tag56.html" target="tagFrame">TargetCompID</a></td>
            <td class="FldNameOdd">
                @TID</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Assigned value used to identify receiving firm.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag57.html" target="tagFrame">57</a></td>
            <td class="FldNameEven"><a href="tag57.html" target="tagFrame">TargetSubID</a></td>
            <td class="FldNameEven">
                @TSub</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Assigned value used to identify specific individual or unit intended to receive message. "ADMIN" reserved for administrative messages not intended for a specific user.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag58.html" target="tagFrame">58</a></td>
            <td class="FldNameOdd"><a href="tag58.html" target="tagFrame">Text</a></td>
            <td class="FldNameOdd">
                @Txt</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Free format text string</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Note: this field does not have a specified maximum length)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag59.html" target="tagFrame">59</a></td>
            <td class="FldNameEven"><a href="tag59.html" target="tagFrame">TimeInForce</a></td>
            <td class="FldNameEven">
                @TmInForce</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Specifies how long the order remains in effect. Absence of this field is interpreted as DAY. NOTE not applicable to CIV Orders. (see Volume : "Glossary" for value definitions)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag60.html" target="tagFrame">60</a></td>
            <td class="FldNameOdd"><a href="tag60.html" target="tagFrame">TransactTime</a></td>
            <td class="FldNameOdd">
                @TxnTm</td>
            <td class="FldDatOdd">UTCTimestamp</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Timestamp when the business transaction represented by the message occurred.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP1 &nbsp;EP94
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag61.html" target="tagFrame">61</a></td>
            <td class="FldNameEven"><a href="tag61.html" target="tagFrame">Urgency</a></td>
            <td class="FldNameEven">
                @Urgency</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Urgency flag</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag62.html" target="tagFrame">62</a></td>
            <td class="FldNameOdd"><a href="tag62.html" target="tagFrame">ValidUntilTime</a></td>
            <td class="FldNameOdd">
                @ValidUntilTm</td>
            <td class="FldDatOdd">UTCTimestamp</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Indicates expiration time of indication message (always expressed in UTC (Universal Time Coordinated, also known as "GMT")</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag63.html" target="tagFrame">63</a></td>
            <td class="FldNameEven"><a href="tag63.html" target="tagFrame">SettlType</a></td>
            <td class="FldNameEven">
                @SettlTyp</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven">Tenor</td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Indicates order settlement period. If present, SettlDate (64) overrides this field. If both SettlType (63) and SettDate (64) are omitted, the default for SettlType (63) is 0 (Regular)</span></p><span class="">
         </span>
                <p class="Even"><span class="">Regular is defined as the default settlement period for the particular security on the exchange of execution.</span></p><span class="">
         </span>
                <p class="Even"><span class="">In Fixed Income the contents of this field may influence the instrument definition if the SecurityID (48) is ambiguous. In the US an active Treasury offering may be re-opened, and for a time one CUSIP will apply to both the current and "when-issued" securities. Supplying a value of "7" clarifies the instrument description; any other value or the absence of this field should cause the respondent to default to the active issue.</span></p><span class="">
         </span>
                <p class="Even"><span class="">Additionally the following patterns may be uses as well as enum values</span></p><span class="">
         </span>
                <p class="Even"><span class="">Dx = FX tenor expression for "days", e.g. "D5", where "x" is any integer &gt; 0</span></p><span class="">
         </span>
                <p class="Even"><span class="">Mx = FX tenor expression for "months", e.g. "M3", where "x" is any integer &gt; 0</span></p><span class="">
         </span>
                <p class="Even"><span class="">Wx = FX tenor expression for "weeks", e.g. "W13", where "x" is any integer &gt; 0</span></p><span class="">
         </span>
                <p class="Even"><span class="">Yx = FX tenor expression for "years", e.g. "Y1", where "x" is any integer &gt; 0</span></p><span class="">
         </span>
                <p class="Even"><span class="">Noted that for FX the tenors expressed using Dx, Mx, Wx, and Yx values do not denote business days, but calendar days.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag64.html" target="tagFrame">64</a></td>
            <td class="FldNameOdd"><a href="tag64.html" target="tagFrame">SettlDate</a></td>
            <td class="FldNameOdd">
                @SettlDt</td>
            <td class="FldDatOdd">LocalMktDate</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Specific date of trade settlement (SettlementDate) in YYYYMMDD format.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">If present, this field overrides SettlType (63). This field is required if the value of SettlType (63) is 6 (Future) or 8 (Sellers Option). This field must be omitted if the value of SettlType (63) is 7 (When and If Issued)</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(expressed in local time at place of settlement)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag65.html" target="tagFrame">65</a></td>
            <td class="FldNameEven"><a href="tag65.html" target="tagFrame">SymbolSfx</a></td>
            <td class="FldNameEven">
                @Sfx</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Additional information about the security (e.g. preferred, warrants, etc.). Note also see SecurityType (167).</span></p><span class="">
         </span>
                <p class="Even"><span class="">As defined in the NYSE Stock and bond Symbol Directory and in the AMEX Fitch Directory.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag66.html" target="tagFrame">66</a></td>
            <td class="FldNameOdd"><a href="tag66.html" target="tagFrame">ListID</a></td>
            <td class="FldNameOdd">
                @ListID
                <br> @ID in ProgramTrading</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Unique identifier for list as assigned by institution, used to associate multiple individual orders. Uniqueness must be guaranteed within a single trading day. Firms which generate multi-day orders should consider embedding a date within the ListID field to assure uniqueness across days.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag67.html" target="tagFrame">67</a></td>
            <td class="FldNameEven"><a href="tag67.html" target="tagFrame">ListSeqNo</a></td>
            <td class="FldNameEven">
                @ListSeqNo
                <br> @SeqNo in ProgramTrading</td>
            <td class="FldDatEven">int</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Sequence of individual order within list (i.e. ListSeqNo of TotNoOrders (68), 2 of 25, 3 of 25, . . . )</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag68.html" target="tagFrame">68</a></td>
            <td class="FldNameOdd"><a href="tag68.html" target="tagFrame">TotNoOrders</a></td>
            <td class="FldNameOdd">
                @TotNoOrds</td>
            <td class="FldDatOdd">int</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Total number of list order entries across all messages. Should be the sum of all NoOrders (73) in each message that has repeating list order entries related to the same ListID (66). Used to support fragmentation.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.2 this field was named "ListNoOrds")</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag69.html" target="tagFrame">69</a></td>
            <td class="FldNameEven"><a href="tag69.html" target="tagFrame">ListExecInst</a></td>
            <td class="FldNameEven">
                @ListExecInst</td>
            <td class="FldDatEven">String</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Free format text message containing list handling and execution instructions.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag70.html" target="tagFrame">70</a></td>
            <td class="FldNameOdd"><a href="tag70.html" target="tagFrame">AllocID</a></td>
            <td class="FldNameOdd">
                @AllocID
                <br> @ID in Allocation</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Unique identifier for allocation message.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag71.html" target="tagFrame">71</a></td>
            <td class="FldNameEven"><a href="tag71.html" target="tagFrame">AllocTransType</a></td>
            <td class="FldNameEven">
                @TransTyp</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Identifies allocation transaction type *** SOME VALUES HAVE BEEN REPLACED - See "Replaced Features and Supported Approach" ***</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag72.html" target="tagFrame">72</a></td>
            <td class="FldNameOdd"><a href="tag72.html" target="tagFrame">RefAllocID</a></td>
            <td class="FldNameOdd">
                @RefAllocID
                <br> @RefID in Allocation</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Reference identifier to be used with AllocTransType (71) = Replace or Cancel.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">(Prior to FIX 4.1 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag73.html" target="tagFrame">73</a></td>
            <td class="FldNameEven"><a href="tag73.html" target="tagFrame">NoOrders</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">NumInGroup</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Indicates number of orders to be combined for average pricing and allocation.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag74.html" target="tagFrame">74</a></td>
            <td class="FldNameOdd"><a href="tag74.html" target="tagFrame">AvgPxPrecision</a></td>
            <td class="FldNameOdd">
                @AvgPxPrcsn</td>
            <td class="FldDatOdd">int</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Indicates number of decimal places to be used for average pricing. Absence of this field indicates that default precision arranged by the broker/institution is to be used.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag75.html" target="tagFrame">75</a></td>
            <td class="FldNameEven"><a href="tag75.html" target="tagFrame">TradeDate</a></td>
            <td class="FldNameEven">
                @TrdDt</td>
            <td class="FldDatEven">LocalMktDate</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Indicates date of trading day. Absence of this field indicates current day (expressed in local time at place of trade).</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP2 &nbsp;EP190
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag77.html" target="tagFrame">77</a></td>
            <td class="FldNameOdd"><a href="tag77.html" target="tagFrame">PositionEffect</a></td>
            <td class="FldNameOdd">
                @PosEfct</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Indicates whether the resulting position after a trade should be an opening position or closing position. Used for omnibus accounting - where accounts are held on a gross basis instead of being netted together.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag78.html" target="tagFrame">78</a></td>
            <td class="FldNameEven"><a href="tag78.html" target="tagFrame">NoAllocs</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">NumInGroup</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Number of repeating AllocAccount (79)/AllocPrice (366) entries.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag79.html" target="tagFrame">79</a></td>
            <td class="FldNameOdd"><a href="tag79.html" target="tagFrame">AllocAccount</a></td>
            <td class="FldNameOdd">
                @Acct</td>
            <td class="FldDatOdd">String</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Sub-account mnemonic</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag80.html" target="tagFrame">80</a></td>
            <td class="FldNameEven"><a href="tag80.html" target="tagFrame">AllocQty</a></td>
            <td class="FldNameEven">
                @Qty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Quantity to be allocated to specific sub-account</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag81.html" target="tagFrame">81</a></td>
            <td class="FldNameOdd"><a href="tag81.html" target="tagFrame">ProcessCode</a></td>
            <td class="FldNameOdd">
                @ProcCode</td>
            <td class="FldDatOdd">char</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Processing code for sub-account. Absence of this field in AllocAccount (79) / AllocPrice (366) /AllocQty (80) / ProcessCode instance indicates regular trade.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag82.html" target="tagFrame">82</a></td>
            <td class="FldNameEven"><a href="tag82.html" target="tagFrame">NoRpts</a></td>
            <td class="FldNameEven">
                @NoRpts</td>
            <td class="FldDatEven">int</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Total number of reports within series.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag83.html" target="tagFrame">83</a></td>
            <td class="FldNameOdd"><a href="tag83.html" target="tagFrame">RptSeq</a></td>
            <td class="FldNameOdd">
                @RptSeq</td>
            <td class="FldDatOdd">int</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Sequence number of message within report series. Used to carry reporting sequence number of the fill as represented on the Trade Report Side.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag84.html" target="tagFrame">84</a></td>
            <td class="FldNameEven"><a href="tag84.html" target="tagFrame">CxlQty</a></td>
            <td class="FldNameEven">
                @CxlQty</td>
            <td class="FldDatEven">Qty</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Total quantity canceled for this order.</span></p><span class="">
         </span>
                <p class="Even"><span class="">(Prior to FIX 4.2 this field was of type int)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag85.html" target="tagFrame">85</a></td>
            <td class="FldNameOdd"><a href="tag85.html" target="tagFrame">NoDlvyInst</a></td>
            <td class="FldNameOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatOdd">NumInGroup</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Number of delivery instruction fields in repeating group.</span></p><span class="">
         </span>
                <p class="Odd"><span class="">Note this field was removed in FIX 4.1 and reinstated in FIX 4.4.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag87.html" target="tagFrame">87</a></td>
            <td class="FldNameEven"><a href="tag87.html" target="tagFrame">AllocStatus</a></td>
            <td class="FldNameEven">
                @Stat
                <br> @Stat in Allocation</td>
            <td class="FldDatEven">int</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Identifies status of allocation.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag88.html" target="tagFrame">88</a></td>
            <td class="FldNameOdd"><a href="tag88.html" target="tagFrame">AllocRejCode</a></td>
            <td class="FldNameOdd">
                @RejCode</td>
            <td class="FldDatOdd">int</td>
            <td class="FldDatOdd">Reserved100Plus</td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Identifies reason for rejection.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
                <br>Updated&nbsp; FIX.5.0SP1 &nbsp;EP95
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagIsDeprEven"><a href="tag89.html" target="tagFrame">89</a></td>
            <td class="FldNameIsDeprEven"><a href="tag89.html" target="tagFrame">Signature</a></td>
            <td class="FldNameIsDeprEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatIsDeprEven">data</td>
            <td class="FldDatIsDeprEven"></td>
            <td class="FldDesIsDeprEven"><span class="">
         </span>
                <p class="IsDeprEven"><span class="">Electronic signature</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprIsDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprIsDeprEven">FIXT.1.1</td>
            <td class="FldEnumIsDeprEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagIsDeprOdd"><a href="tag90.html" target="tagFrame">90</a></td>
            <td class="FldNameIsDeprOdd"><a href="tag90.html" target="tagFrame">SecureDataLen</a></td>
            <td class="FldNameIsDeprOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatIsDeprOdd">Length</td>
            <td class="FldDatIsDeprOdd"></td>
            <td class="FldDesIsDeprOdd"><span class="">
         </span>
                <p class="IsDeprOdd"><span class="">Length of encrypted message</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprIsDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprIsDeprOdd">FIXT.1.1</td>
            <td class="FldEnumIsDeprOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagIsDeprEven"><a href="tag91.html" target="tagFrame">91</a></td>
            <td class="FldNameIsDeprEven"><a href="tag91.html" target="tagFrame">SecureData</a></td>
            <td class="FldNameIsDeprEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatIsDeprEven">data</td>
            <td class="FldDatIsDeprEven"></td>
            <td class="FldDesIsDeprEven"><span class="">
         </span>
                <p class="IsDeprEven"><span class="">Actual encrypted data stream</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprIsDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprIsDeprEven">FIXT.1.1</td>
            <td class="FldEnumIsDeprEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagIsDeprOdd"><a href="tag93.html" target="tagFrame">93</a></td>
            <td class="FldNameIsDeprOdd"><a href="tag93.html" target="tagFrame">SignatureLength</a></td>
            <td class="FldNameIsDeprOdd"><em>(not used in FIXML)</em></td>
            <td class="FldDatIsDeprOdd">Length</td>
            <td class="FldDatIsDeprOdd"></td>
            <td class="FldDesIsDeprOdd"><span class="">
         </span>
                <p class="IsDeprOdd"><span class="">Number of bytes in signature field</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprIsDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprIsDeprOdd">FIXT.1.1</td>
            <td class="FldEnumIsDeprOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag94.html" target="tagFrame">94</a></td>
            <td class="FldNameEven"><a href="tag94.html" target="tagFrame">EmailType</a></td>
            <td class="FldNameEven">
                @EmailTyp</td>
            <td class="FldDatEven">char</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Email message type.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag95.html" target="tagFrame">95</a></td>
            <td class="FldNameOdd"><a href="tag95.html" target="tagFrame">RawDataLength</a></td>
            <td class="FldNameOdd">
                @RawDataLength</td>
            <td class="FldDatOdd">Length</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Number of bytes in raw data field.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag96.html" target="tagFrame">96</a></td>
            <td class="FldNameEven"><a href="tag96.html" target="tagFrame">RawData</a></td>
            <td class="FldNameEven">
                @RawData</td>
            <td class="FldDatEven">data</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Unformatted raw data, can include bitmaps, word processor documents, etc.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag97.html" target="tagFrame">97</a></td>
            <td class="FldNameOdd"><a href="tag97.html" target="tagFrame">PossResend</a></td>
            <td class="FldNameOdd">
                @PosRsnd</td>
            <td class="FldDatOdd">Boolean</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Indicates that message may contain information that has been sent under another sequence number.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag98.html" target="tagFrame">98</a></td>
            <td class="FldNameEven"><a href="tag98.html" target="tagFrame">EncryptMethod</a></td>
            <td class="FldNameEven"><em>(not used in FIXML)</em></td>
            <td class="FldDatEven">int</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Method of encryption.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag99.html" target="tagFrame">99</a></td>
            <td class="FldNameOdd"><a href="tag99.html" target="tagFrame">StopPx</a></td>
            <td class="FldNameOdd">
                @StopPx</td>
            <td class="FldDatOdd">Price</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Price per unit of quantity (e.g. per share)</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag100.html" target="tagFrame">100</a></td>
            <td class="FldNameEven"><a href="tag100.html" target="tagFrame">ExDestination</a></td>
            <td class="FldNameEven">
                @ExDest</td>
            <td class="FldDatEven">Exchange</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Execution destination as defined by institution when order is entered.</span></p><span class="">
         </span>
                <p class="Even"><span class="">Valid values:</span></p><span class="">
         </span>
                <p class="Even"><span class="">See "Appendix 6-C"</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.2.7
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        
        <tr xmlns="">
            <td class="FldTagOdd"><a href="tag50001.html" target="tagFrame">50001</a></td>
            <td class="FldNameOdd"><a href="tag50001.html" target="tagFrame">BatchTotalMessages</a></td>
            <td class="FldNameOdd">
                @TotMsg</td>
            <td class="FldDatOdd">int</td>
            <td class="FldDatOdd"></td>
            <td class="FldDesOdd"><span class="">
         </span>
                <p class="Odd"><span class="">Total # of messages contained within batch.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprOdd">
                Added&nbsp; FIX.5.0SP2 &nbsp;EP178
            </td>
            <td class="FldDeprOdd"></td>
            <td class="FldEnumOdd">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
        <tr xmlns="">
            <td class="FldTagEven"><a href="tag50002.html" target="tagFrame">50002</a></td>
            <td class="FldNameEven"><a href="tag50002.html" target="tagFrame">BatchProcessMode</a></td>
            <td class="FldNameEven">
                @ProcMode</td>
            <td class="FldDatEven">int</td>
            <td class="FldDatEven"></td>
            <td class="FldDesEven"><span class="">
         </span>
                <p class="Even"><span class="">Indicates the processing mode for a batch of messages.</span></p><span class="">
      </span>
                <br>
                <br>
            </td>
            <td class="FldDeprEven">
                Added&nbsp; FIX.5.0SP2 &nbsp;EP178
            </td>
            <td class="FldDeprEven"></td>
            <td class="FldEnumEven">
                <a href="tag.html" target="tagFrame"></a>
            </td>
        </tr>
    </table>
    <p xmlns=""></p>
    <p xmlns=""></p>
    <table xmlns="">
        <tr>
            <td class="Horiz" id="footer" colspan="2">
                <ul class="Menu">
                    <li id="copyright">© 2007 - 2017 FIX Protocol Limited</li>
                    <li><a href="http://www.fixtrading.org/contact-us" target="_blank">Contact us</a></li>
                    <li><a href="http://www.fixtrading.org/terms-and-conditions" target="_blank">Copyright and Acceptable Use policy</a></li>
                    <li><a href="http://www.fixtrading.org/privacy-policy" target="_blank">Privacy Policy</a></li>
                </ul>
            </td>
        </tr>
    </table>
</body>

</html>`

const html1 = `<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
    <head>
        <link rel="stylesheet" href="../../fixstyle.css" type="text/css"></link>
        <title>FIX.5.0SP2_EP254 Field #35</title>
        <meta http-equiv="Content-Type"
              content="text/html; charset=utf-8"></meta>
    </head>
    <body>
        <table>
            <tr>
                <th class="FldTagHdr">Tag</th>
                <th class="FldNameHdr">Field Name</th>
                <th class="FldNameHdr">XML Name</th>
                <th class="FldDatHdr">Data Type</th>
                <th class="FldDatHdr">Union Datatype</th>
                <th class="FldDesHdr">Description</th>
                <th class="FldDesHdr">Added</th>
                <th class="FldDesHdr">Depr.</th>
                <th class="FldDesHdr">Enums from tag</th>
            </tr>
            <tr xmlns="">
                <td class="FldTagEven"><a href="tag35.html"
                                          target="_blank">35</a></td>
                <td class="FldNameEven"><a href="tag35.html" target="_blank">MsgType</a>
                </td>
                <td class="FldNameEven">
                    @MsgTyp
                </td>
                <td class="FldDatEven">String</td>
                <td class="FldDatEven"></td>
                <td class="FldDesEven"><span class="">
         </span>
                    <p class="Even"><span class="">Defines message type ALWAYS THIRD FIELD IN MESSAGE. (Always unencrypted)</span>
                    </p><span class="">
         </span>
                    <p class="Even"><span class="">Note: A "U" as the first character in the MsgType field (i.e. U, U2, etc) indicates that the message format is privately defined between the sender and receiver.</span>
                    </p><span class="">
         </span>
                    <p class="Even"><span class="">*** Note the use of lower case letters ***</span>
                    </p><span class="">
      </span><br><br>
                    <table>
                        <tr>
                            <td class="EnmTwo">0</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">Heartbeat</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Heartbeat monitors the status of the communication link and identifies when the last of a string of messages was not received.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[Heartbeat]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">1</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">TestRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The test request message forces a heartbeat from the opposing application. The test request message checks sequence numbers or verifies communication line status. The opposite application responds to the Test Request with a Heartbeat containing the TestReqID.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TestRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">2</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">ResendRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The resend request is sent by the receiving application to initiate the retransmission of messages. This function is utilized if a sequence number gap is detected, if the receiving application lost a message, or as a function of the initialization process.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ResendRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">3</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">Reject</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The reject message should be issued when a message is received but cannot be properly processed due to a session-level rule violation. An example of when a reject may be appropriate would be the receipt of a message with invalid basic data which successfully passes de-encryption, CheckSum and BodyLength checks.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[Reject]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">4</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">SequenceReset</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The sequence reset message is used by the sending application to reset the incoming sequence number on the opposing side.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SequenceReset]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">5</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">Logout</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The logout message initiates or confirms the termination of a FIX session. Disconnection without the exchange of logout messages should be interpreted as an abnormal condition.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[Logout]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">6</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">IOI</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Indication of interest messages are used to market merchandise which the broker is buying or selling in either a proprietary or agency capacity. The indications can be time bound with a specific expiration value. Indications are distributed with the understanding that other firms may react to the message first and that the merchandise may no longer be available due to prior trade.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">Indication messages can be transmitted in various transaction types; NEW, CANCEL, and REPLACE. All message types other than NEW modify the state of the message identified in IOIRefID.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[IOI]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">7</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">Advertisement</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Advertisement messages are used to announce completed transactions. The advertisement message can be transmitted in various transaction types; NEW, CANCEL and REPLACE. All message types other than NEW modify the state of a previously transmitted advertisement identified in AdvRefID.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[Advertisement]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">8</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">ExecutionReport</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The execution report message is used to:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">1. confirm the receipt of an order</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">2. confirm changes to an existing order (i.e. accept cancel and replace requests)</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">3. relay order status information</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">4. relay fill information on working orders</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">5. relay fill information on tradeable or restricted tradeable quotes</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">6. reject orders</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">7. report post-trade fees calculations associated with a trade</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ExecutionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">9</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">OrderCancelReject</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order cancel reject message is issued by the broker upon receipt of a cancel request or cancel/replace request message which cannot be honored.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[OrderCancelReject]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">A</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">Logon</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The logon message authenticates a user establishing a connection to a remote system. The logon message must be the first message sent by the application requesting to initiate a FIX session.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[Logon]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">B</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">News</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The news message is a general free format message between the broker and institution. The message contains flags to identify the news item's urgency and to allow sorting by subject company (symbol). The News message can be originated at either the broker or institution side, or exchanges and other marketplace venues.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[News]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">C</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">Email</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The email message is similar to the format and purpose of the News message, however, it is intended for private use between two parties.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[Email]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">D</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">NewOrderSingle</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The new order message type is used by institutions wishing to electronically submit securities and forex orders to a broker for execution.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">The New Order message type may also be used by institutions or retail intermediaries wishing to electronically submit Collective Investment Vehicle (CIV) orders to a broker or fund manager for execution.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[NewOrderSingle]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">E</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">NewOrderList</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The NewOrderList Message can be used in one of two ways depending on which market conventions are being followed.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[NewOrderList]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">F</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">OrderCancelRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order cancel request message requests the cancellation of all of the remaining quantity of an existing order. Note that the Order Cancel/Replace Request should be used to partially cancel (reduce) an order).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[OrderCancelRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">G</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">OrderCancelReplaceRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order cancel/replace request is used to change the parameters of an existing order.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">Do not use this message to cancel the remaining quantity of an outstanding order, use the Order Cancel Request message for this purpose.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[OrderCancelReplaceRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">H</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">OrderStatusRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order status request message is used by the institution to generate an order status message back from the broker.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[OrderStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">J</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">AllocationInstruction</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Allocation Instruction message provides the ability to specify how an order or set of orders should be subdivided amongst one or more accounts. In versions of FIX prior to version 4.4, this same message was known as the Allocation message. Note in versions of FIX prior to version 4.4, the allocation message was also used to communicate fee and expense details from the Sellside to the Buyside. This role has now been removed from the Allocation Instruction and is now performed by the new (to version 4.4) Allocation Report and Confirmation messages.,The Allocation Report message should be used for the Sell-side Initiated Allocation role as defined in previous versions of the protocol.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[AllocationInstruction]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">K</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">ListCancelRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The List Cancel Request message type is used by institutions wishing to cancel previously submitted lists either before or during execution.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ListCancelRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">L</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">ListExecute</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The List Execute message type is used by institutions to instruct the broker to begin execution of a previously submitted list. This message may or may not be used, as it may be mirroring a phone conversation.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ListExecute]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">M</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">ListStatusRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The list status request message type is used by institutions to instruct the broker to generate status messages for a list.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ListStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">N</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">ListStatus</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The list status message is issued as the response to a List Status Request message sent in an unsolicited fashion by the sell-side. It indicates the current state of the orders within the list as they exist at the broker's site. This message may also be used to respond to the List Cancel Request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ListStatus]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">P</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">AllocationInstructionAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">In versions of FIX prior to version 4.4, this message was known as the Allocation ACK message.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">The Allocation Instruction Ack message is used to acknowledge the receipt of and provide status for an Allocation Instruction message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[AllocationInstructionAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">Q</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">DontKnowTrade</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Don’t Know Trade (DK) message notifies a trading partner that an electronically received execution has been rejected. This message can be thought of as an execution reject message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[DontKnowTrade]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">R</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">QuoteRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">In some markets it is the practice to request quotes from brokers prior to placement of an order. The quote request message is used for this purpose. This message is commonly referred to as an Request For Quote (RFQ)</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[QuoteRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">S</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">Quote</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Quote message is used as the response to a Quote Request or a Quote Response message in both indicative, tradeable, and restricted tradeable quoting markets.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[Quote]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">T</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SettlementInstructions</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Settlement Instructions message provides the broker’s, the institution’s, or the intermediary’s instructions for trade settlement. This message has been designed so that it can be sent from the broker to the institution, from the institution to the broker, or from either to an independent "standing instructions" database or matching system or, for CIV, from an intermediary to a fund manager.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SettlementInstructions]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">V</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">MarketDataRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Some systems allow the transmission of real-time quote, order, trade, trade volume, open interest, and/or other price information on a subscription basis. A MarketDataRequest(35=V) is a general request for market data on specific securities or forex quotes. The values in the fields provided within the request will serve as further filter criteria for the result set.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarketDataRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">W</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MarketDataSnapshotFullRefresh</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Market Data messages are used as the response to a Market Data Request message. In all cases, one Market Data message refers only to one Market Data Request. It can be used to transmit a 2-sided book of orders or list of quotes, a list of trades, index values, opening, closing, settlement, high, low, or VWAP prices, the trade volume or open interest for a security, or any combination of these.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarketDataSnapshotFullRefresh]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">X</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarketDataIncrementalRefresh</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Market Data message for incremental updates may contain any combination of new, changed, or deleted Market Data Entries, for any combination of instruments, with any combination of trades, imbalances, quotes, index values, open, close, settlement, high, low, and VWAP prices, trade volume and open interest so long as the maximum FIX message size is not exceeded. All of these types of Market Data Entries can be changed and deleted.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarketDataIncrementalRefresh]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">Y</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MarketDataRequestReject</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Market Data Request Reject is used when the broker cannot honor the Market Data Request, due to business or technical reasons. Brokers may choose to limit various parameters, such as the size of requests, whether just the top of book or the entire book may be displayed, and whether Full or Incremental updates must be used.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarketDataRequestReject]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">Z</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">QuoteCancel</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Quote Cancel message is used by an originator of quotes to cancel quotes.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">The Quote Cancel message supports cancellation of:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span
                                            class="">• All quotes</span></p>
                                    <span class="">
         </span>
                                    <p class="Even"><span class="">• Quotes for a specific symbol or security ID</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• All quotes for a security type</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• All quotes for an underlying</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[QuoteCancel]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">a</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">QuoteStatusRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The quote status request message is used for the following purposes in markets that employ tradeable or restricted tradeable quotes:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• For the issuer of a quote in a market to query the status of that quote (using the QuoteID to specify the target quote).</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• To subscribe and unsubscribe for Quote Status Report messages for one or more securities.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[QuoteStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">b</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MassQuoteAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Mass Quote Acknowledgement is used as the application level response to a Mass Quote message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MassQuoteAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">c</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SecurityDefinitionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The SecurityDefinitionRequest(35=c) message is used for the following:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">1. Request a specific security to be traded with the second party. The requested security can be defined as a multileg security made up of one or more instrument legs.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">2. Request a set of individual securities for a single market segment.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">3. Request all securities, independent of market segment.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityDefinitionRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">d</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">SecurityDefinition</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The SecurityDefinition(35=d) message is used for the following:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">1. Accept the security defined in a SecurityDefinition(35=d) message.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">2. Accept the security defined in a SecurityDefinition(35=d) message with changes to the definition and/or identity of the security.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">3. Reject the security requested in a SecurityDefinition(35=d) message.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">4. Respond to a request for securities within a specified market segment.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">5. Convey comprehensive security definition for all market segments that the security participates in.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">6. Convey the security's trading rules that differ from default rules for the market segment.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SecurityDefinition]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">e</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">SecurityStatusRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security Status Request message provides for the ability to request the status of a security. One or more Security Status messages are returned as a result of a Security Status Request message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">f</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">SecurityStatus</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security Status message provides for the ability to report changes in status to a security. The Security Status message contains fields to indicate trading status, corporate actions, financial status of the company. The Security Status message is used by one trading entity (for instance an exchange) to report changes in the state of a security.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SecurityStatus]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">g</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">TradingSessionStatusRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trading Session Status Request is used to request information on the status of a market. With the move to multiple sessions occurring for a given trading party (morning and evening sessions for instance) there is a need to be able to provide information on what product is trading on what market.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradingSessionStatusRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">h</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradingSessionStatus</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trading Session Status provides information on the status of a market. For markets multiple trading sessions on multiple-markets occurring (morning and evening sessions for instance), this message is able to provide information on what products are trading on what market during what trading session.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradingSessionStatus]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">i</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MassQuote</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Mass Quote message can contain quotes for multiple securities to support applications that allow for the mass quoting of an option series. Two levels of repeating groups have been provided to minimize the amount of data required to submit a set of quotes for a class of options (e.g. all option series for IBM).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MassQuote]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">j</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">BusinessMessageReject</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Business Message Reject message can reject an application-level message which fulfills session-level rules and cannot be rejected via any other means. Note if the message fails a session-level rule (e.g. body length is incorrect), a session-level Reject message should be issued.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[BusinessMessageReject]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">k</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">BidRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The BidRequest Message can be used in one of two ways depending on which market conventions are being followed.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">In the "Non disclosed" convention (e.g. US/European model) the BidRequest message can be used to request a bid based on the sector, country, index and liquidity information contained within the message itself. In the "Non disclosed" convention the entry repeating group is used to define liquidity of the program. See " Program/Basket/List Trading" for an example.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">In the "Disclosed" convention (e.g. Japanese model) the BidRequest message can be used to request bids based on the ListOrderDetail messages sent in advance of BidRequest message. In the "Disclosed" convention the list repeating group is used to define which ListOrderDetail messages a bid is being sort for and the directions of the required bids.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[BidRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">l</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">BidResponse</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Bid Response message can be used in one of two ways depending on which market conventions are being followed.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">In the "Non disclosed" convention the Bid Response message can be used to supply a bid based on the sector, country, index and liquidity information contained within the corresponding bid request message. See "Program/Basket/List Trading" for an example.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">In the "Disclosed" convention the Bid Response message can be used to supply bids based on the List Order Detail messages sent in advance of the corresponding Bid Request message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[BidResponse]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">m</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">ListStrikePrice</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The strike price message is used to exchange strike price information for principal trades. It can also be used to exchange reference prices for agency trades.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ListStrikePrice]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">n</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">XMLnonFIX</span></p>
                                <span class="">
      </span><i></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[XMLnonFIX]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">o</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">RegistrationInstructions</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Registration Instructions message type may be used by institutions or retail intermediaries wishing to electronically submit registration information to a broker or fund manager (for CIV) for an order or for an allocation.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[RegistrationInstructions]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">p</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">RegistrationInstructionsResponse</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Registration Instructions Response message type may be used by broker or fund manager (for CIV) in response to a Registration Instructions message submitted by an institution or retail intermediary for an order or for an allocation.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [RegistrationInstructionsResponse]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">q</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">OrderMassCancelRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order mass cancel request message requests the cancellation of all of the remaining quantity of a group of orders matching criteria specified within the request. NOTE: This message can only be used to cancel order messages (reduce the full quantity).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[OrderMassCancelRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">r</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">OrderMassCancelReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Order Mass Cancel Report is used to acknowledge an Order Mass Cancel Request. Note that each affected order that is canceled is acknowledged with a separate Execution Report or Order Cancel Reject message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[OrderMassCancelReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">s</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">NewOrderCross</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to submit a cross order into a market. The cross order contains two order sides (a buy and a sell). The cross order is identified by its CrossID.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[NewOrderCross]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">t</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">CrossOrderCancelReplaceRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to modify a cross order previously submitted using the New Order - Cross message. See Order Cancel Replace Request for details concerning message usage.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [CrossOrderCancelReplaceRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">u</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">CrossOrderCancelRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to fully cancel the remaining open quantity of a cross order.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[CrossOrderCancelRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">v</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">SecurityTypeRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security Type Request message is used to return a list of security types available from a counterparty or market.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SecurityTypeRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">w</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SecurityTypes</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security Type Request message is used to return a list of security types available from a counterparty or market.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityTypes]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">x</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">SecurityListRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security List Request message is used to return a list of securities from the counterparty that match criteria provided on the request</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SecurityListRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">y</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SecurityList</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security List message is used to return a list of securities that matches the criteria specified in a Security List Request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityList]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">z</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">DerivativeSecurityListRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Derivative Security List Request message is used to return a list of securities from the counterparty that match criteria provided on the request</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[DerivativeSecurityListRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AA</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">DerivativeSecurityList</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Derivative Security List message is used to return a list of securities that matches the criteria specified in a Derivative Security List Request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[DerivativeSecurityList]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AB</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">NewOrderMultileg</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The New Order - Multileg is provided to submit orders for securities that are made up of multiple securities, known as legs.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[NewOrderMultileg]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AC</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MultilegOrderCancelReplace</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to modify a multileg order previously submitted using the New Order - Multileg message. See Order Cancel Replace Request for details concerning message usage.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MultilegOrderCancelReplace]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AD</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">TradeCaptureReportRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trade Capture Report Request can be used to:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• Request one or more trade capture reports based upon selection criteria provided on the trade capture report request</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• Subscribe for trade capture reports based upon selection criteria provided on the trade capture report request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradeCaptureReportRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AE</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradeCaptureReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trade Capture Report message can be:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Used to report trades between counterparties.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Used to report trades to a trade matching system.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Sent unsolicited between counterparties.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Sent as a reply to a Trade Capture Report Request.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Used to report unmatched and matched trades.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradeCaptureReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AF</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">OrderMassStatusRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The order mass status request message requests the status for orders matching criteria specified within the request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[OrderMassStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AG</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">QuoteRequestReject</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Quote Request Reject message is used to reject Quote Request messages for all quoting models.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[QuoteRequestReject]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AH</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">RFQRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">In tradeable and restricted tradeable quoting markets – Quote Requests are issued by counterparties interested in ascertaining the market for an instrument. Quote Requests are then distributed by the market to liquidity providers who make markets in the instrument. The RFQ Request is used by liquidity providers to indicate to the market for which instruments they are interested in receiving Quote Requests. It can be used to register interest in receiving quote requests for a single instrument or for multiple instruments</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[RFQRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AI</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">QuoteStatusReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The quote status report message is used:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• as the response to a Quote Status Request message</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• as a response to a Quote Cancel message</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• as a response to a Quote Response message in a negotiation dialog (see Volume 7 – PRODUCT: FIXED INCOME and USER GROUP: EXCHANGES AND MARKETS)</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[QuoteStatusReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AJ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">QuoteResponse</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The QuoteResponse(35=AJ) message is used for the following purposes:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">1. Respond to an IOI(35=6) message</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">2. Respond to a Quote(35=S) message</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">3. Counter a Quote</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">4. End a negotiation dialog</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">5. Follow-up or end a QuoteRequest(35=R) dialog that did not receive a response.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[QuoteResponse]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AK</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">Confirmation</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Confirmation messages are used to provide individual trade level confirmations from the sell side to the buy side. In versions of FIX prior to version 4.4, this role was performed by the allocation message. Unlike the allocation message, the confirmation message operates at an allocation account (trade) level rather than block level, allowing for the affirmation or rejection of individual confirmations.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[Confirmation]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AL</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PositionMaintenanceRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Position Maintenance Request message allows the position owner to submit requests to the holder of a position which will result in a specific action being taken which will affect the position. Generally, the holder of the position is a central counter party or clearing organization but can also be a party providing investment services.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PositionMaintenanceRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AM</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PositionMaintenanceReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Position Maintenance Report message is sent by the holder of a positon in response to a Position Maintenance Request and is used to confirm that a request has been successfully processed or rejected.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PositionMaintenanceReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AN</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">RequestForPositions</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Request For Positions message is used by the owner of a position to request a Position Report from the holder of the position, usually the central counter party or clearing organization. The request can be made at several levels of granularity.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[RequestForPositions]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AO</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">RequestForPositionsAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Request for Positions Ack message is returned by the holder of the position in response to a Request for Positions message. The purpose of the message is to acknowledge that a request has been received and is being processed.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[RequestForPositionsAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AP</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PositionReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Position Report message is returned by the holder of a position in response to a Request for Position message. The purpose of the message is to report all aspects of a position and may be provided on a standing basis to report end of day positions to an owner.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PositionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AQ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">TradeCaptureReportRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trade Capture Request Ack message is used to:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Provide an acknowledgement to a Trade Capture Report Request in the case where the Trade Capture Report Request is used to specify a subscription or delivery of reports via an out-of-band ResponseTransmissionMethod.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Provide an acknowledgement to a Trade Capture Report Request in the case when the return of the Trade Capture Reports matching that request will be delayed or delivered asynchronously. This is useful in distributed trading system environments.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Indicate that no trades were found that matched the selection criteria specified on the Trade Capture Report Request or the Trade Capture Request was invalid for some business reason, such as request is not authorized, invalid or unknown instrument, party, trading session, etc.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradeCaptureReportRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AR</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradeCaptureReportAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trade Capture Report Ack message can be:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Used to acknowledge trade capture reports received from a counterparty.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">- Used to reject a trade capture report received from a counterparty.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradeCaptureReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AS</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">AllocationReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Sent from sell-side to buy-side, sell-side to 3rd-party or 3rd-party to buy-side, the Allocation Report (Claim) provides account breakdown of an order or set of orders plus any additional follow-up front-office information developed post-trade during the trade allocation, matching and calculation phase. In versions of FIX prior to version 4.4, this functionality was provided through the Allocation message. Depending on the needs of the market and the timing of "confirmed" status, the role of Allocation Report can be taken over in whole or in part by the Confirmation message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[AllocationReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AT</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">AllocationReportAck</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Allocation Report Ack message is used to acknowledge the receipt of and provide status for an Allocation Report message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[AllocationReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AU</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">ConfirmationAck</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Confirmation Ack (aka Affirmation) message is used to respond to a Confirmation message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ConfirmationAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AV</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">SettlementInstructionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Settlement Instruction Request message is used to request standing settlement instructions from another party.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SettlementInstructionRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AW</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">AssignmentReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Assignment Reports are sent from a clearing house to counterparties, such as a clearing firm as a result of the assignment process.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[AssignmentReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AX</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">An initiator that requires collateral from a respondent sends a Collateral Request. The initiator can be either counterparty to a trade in a two party model or an intermediary such as an ATS or clearinghouse in a three party model. A Collateral Assignment is expected as a response to a request for collateral.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[CollateralRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">AY</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralAssignment</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to assign collateral to cover a trading position. This message can be sent unsolicited or in reply to a Collateral Request message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[CollateralAssignment]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">AZ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralResponse</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to respond to a Collateral Assignment message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[CollateralResponse]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BA</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to report collateral status when responding to a Collateral Inquiry message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[CollateralReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BB</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralInquiry</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to inquire for collateral status.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[CollateralInquiry]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BC</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">NetworkCounterpartySystemStatusRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is send either immediately after logging on to inform a network (counterparty system) of the type of updates required or to at any other time in the FIX conversation to change the nature of the types of status updates required. It can also be used with a NetworkRequestType of Snapshot to request a one-off report of the status of a network (or counterparty) system. Finally this message can also be used to cancel a request to receive updates into the status of the counterparties on a network by sending a NetworkRequestStatusMessage with a NetworkRequestType of StopSubscribing.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">
                                [NetworkCounterpartySystemStatusRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BD</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">NetworkCounterpartySystemStatusResponse</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is sent in response to a Network (Counterparty System) Status Request Message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [NetworkCounterpartySystemStatusResponse]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BE</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">UserRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used to initiate a user action, logon, logout or password change. It can also be used to request a report on a user's status.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[UserRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BF</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">UserResponse</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used to respond to a user request message, it reports the status of the user after the completion of any action requested in the user request message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[UserResponse]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BG</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralInquiryAck</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to respond to a Collateral Inquiry in the following situations:</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• When the CollateralInquiry will result in an out of band response (such as a file transfer).</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• When the inquiry is otherwise valid but no collateral is found to match the criteria specified on the Collateral Inquiry message.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">• When the Collateral Inquiry is invalid based upon the business rules of the counterparty.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[CollateralInquiryAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BH</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">ConfirmationRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Confirmation Request message is used to request a Confirmation message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ConfirmationRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BO</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">ContraryIntentionReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Contrary Intention Report is used for reporting of contrary expiration quantities for Saturday expiring options. This information is required by options exchanges for regulatory purposes.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ContraryIntentionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BP</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">SecurityDefinitionUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used for reporting updates to a product security master file. Updates could be the result of corporate actions or other business events. Updates may include additions, modifications or deletions.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [SecurityDefinitionUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BK</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SecurityListUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Security List Update Report is used for reporting updates to a Contract Security Masterfile. Updates could be due to Corporate Actions or other business events. Update may include additions, modifications and deletions.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityListUpdateReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BL</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">AdjustedPositionReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to report changes in position, primarily in equity options, due to modifications to the underlying due to corporate actions</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[AdjustedPositionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BM</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">AllocationInstructionAlert</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used in a 3-party allocation model where notification of group creation and group updates to counterparties is needed. The mssage will also carry trade information that comprised the group to the counterparties.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[AllocationInstructionAlert]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BN</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">ExecutionAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Execution Report Acknowledgement message is an optional message that provides dual functionality to notify a trading partner that an electronically received execution has either been accepted or rejected (DK'd).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ExecutionAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BJ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradingSessionList</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trading Session List message is sent as a response to a Trading Session List Request. The Trading Session List should contain the characteristics of the trading session and the current state of the trading session.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradingSessionList]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BI</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">TradingSessionListRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trading Session List Request is used to request a list of trading sessions available in a market place and the state of those trading sessions. A successful request will result in a response from the counterparty of a Trading Session List (MsgType=BJ) message that contains a list of zero or more trading sessions.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradingSessionListRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BQ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">SettlementObligationReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Settlement Obligation Report message provides a central counterparty, institution, or individual counterparty with a capacity for reporting the final details of a currency settlement obligation.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SettlementObligationReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BR</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">DerivativeSecurityListUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Derivative Security List Update Report message is used to send updates to an option family or the strikes that comprise an option family.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [DerivativeSecurityListUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BS</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">TradingSessionListUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Trading Session List Update Report is used by marketplaces to provide intra-day updates of trading sessions when there are changes to one or more trading sessions.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">
                                [TradingSessionListUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BT</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarketDefinitionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Market Definition Request message is used to request for market structure information from the Respondent that receives this request.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarketDefinitionRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BU</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">MarketDefinition</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The MarketDefinition(35=BU) message is used to respond to MarketDefinitionRequest(35=BT). In a subscription, it will be used to provide the initial snapshot of the information requested. Subsequent updates are provided by the MarketDefinitionUpdateReport(35=BV).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarketDefinition]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BV</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarketDefinitionUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">In a subscription for market structure information, this message is used once the initial snapshot of the information has been sent using the MarketDefinition(35=BU) message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarketDefinitionUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BW</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">ApplicationMessageRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used to request a retransmission of a set of one or more messages generated by the application specified in RefApplID (1355).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ApplicationMessageRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BX</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">ApplicationMessageRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used to acknowledge an Application Message Request providing a status on the request (i.e. whether successful or not). This message does not provide the actual content of the messages to be resent.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[ApplicationMessageRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">BY</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">ApplicationMessageReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used for three difference purposes: to reset the ApplSeqNum (1181) of a specified ApplID (1180). to indicate that the last message has been sent for a particular ApplID, or as a keep-alive mechanism for ApplIDs with infrequent message traffic.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[ApplicationMessageReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">BZ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">OrderMassActionReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Order Mass Action Report is used to acknowledge an Order Mass Action Request. Note that each affected order that is suspended or released or canceled is acknowledged with a separate Execution Report for each order.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[OrderMassActionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CA</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">OrderMassActionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Order Mass Action Request message can be used to request the suspension or release of a group of orders that match the criteria specified within the request. This is equivalent to individual Order Cancel Replace Requests for each order with or without adding "S" to the ExecInst values. It can also be used for mass order cancellation.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[OrderMassActionRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CB</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">UserNotification</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The User Notification message is used to notify one or more users of an event or information from the sender of the message. This message is usually sent unsolicited from a marketplace (e.g. Exchange, ECN) to a market participant.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[UserNotification]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CC</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">StreamAssignmentRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">In certain markets where market data aggregators fan out to end clients the pricing streams provided by the price makers, the price maker may assign the clients to certain pricing streams that the price maker publishes via the aggregator. An example of this use is in the FX markets where clients may be assigned to different pricing streams based on volume bands and currency pairs.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[StreamAssignmentRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CD</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">StreamAssignmentReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">he StreamAssignmentReport message is in response to the StreamAssignmentRequest message. It provides information back to the aggregator as to which clients to assign to receive which price stream based on requested CCY pair. This message can be sent unsolicited to the Aggregator from the Price Maker.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[StreamAssignmentReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CE</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">StreamAssignmentReportACK</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used to respond to the Stream Assignment Report, to either accept or reject an unsolicited assingment.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[StreamAssignmentReportACK]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CF</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyDetailsListRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyDetailsListRequest is used to request party detail information.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyDetailsListRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CG</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyDetailsListReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyDetailsListReport message is used to disseminate party details between counterparties. PartyDetailsListReport messages may be sent in response to a PartyDetailsListRequest message or sent unsolicited.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyDetailsListReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CH</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarginRequirementInquiry</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The purpose of this message is to initiate a margin requirement inquiry for a margin account. The inquiry may be submitted at the detail level or the summary level. It can also be used to inquire margin excess/deficit or net position information. Margin excess/deficit will provide information about the surplus or shortfall compared to the previous trading day or a more recent margin calculation. An inquiry for net position information will trigger one or more PositionReport messages instead of one or more MarginRequirementReport messages.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">If the inquiry is made at the detail level, an Instrument block must be provided with the desired level of detail. If the inquiry is made at the summary level, the Instrument block is not provided, implying a summary request is being made. For example, if the inquiring firm specifies the Security Type of “FUT” in the Instrument block, then a detail report will be generated containing the margin requirements for all futures positions for the inquiring account. Similarly, if the inquiry is made at the summary level, the report will contain the total margin requirement aggregated to the margin account level.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarginRequirementInquiry]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CI</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MarginRequirementInquiryAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to respond to a Margin Requirement Inquiry.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarginRequirementInquiryAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CJ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarginRequirementReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The Margin Requirement Report returns information about margin requirement either as on overview across all margin accounts or on a detailed level due to the inquiry making use of the optional Instrument component block. Application sequencing can be used to re-request a range of reports.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarginRequirementReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CK</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyDetailsListUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyDetailsListUpdateReport(35=CK) is used to disseminate updates to party detail information.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyDetailsListUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CL</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitsRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyRiskLimitsRequest message is used to request for risk information for specific parties, specific party roles or specific instruments.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyRiskLimitsRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CM</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">PartyRiskLimitsReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyRiskLimitsReport message is used to communicate party risk limits. The message can either be sent as a response to the PartyRiskLimitsRequest message or can be published unsolicited.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyRiskLimitsReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CN</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">SecurityMassStatusRequest</span>
                                </p><span class="">
      </span><i></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[SecurityMassStatusRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CO</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">SecurityMassStatus</span></p>
                                <span class="">
      </span><i></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[SecurityMassStatus]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CQ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">AccountSummaryReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The AccountSummaryReport is provided by the clearinghouse to its clearing members on a daily basis. It contains margin, settlement, collateral and pay/collect data for each clearing member level account type. Clearing member account types will be described through use of the Parties component and PtysSubGrp sub-component.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">In certain usages, the clearing members can send the AccountSummaryReport message to the clearinghouse as needed. For example, clearing members can send this message to the clearinghouse to identify the value of collateral for each customer (to satisfy CFTC Legally Segregated Operationally Commingled (LSOC) regulatory reporting obligations).</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">Clearing organizations can also send the AccountSummaryReport message to regulators to meet regulatory reporting obligations. For example, clearing organizations can use this message to submit daily reports for each clearing member (“CM”) by house origin and by each customer origin for all futures, options, and swaps positions, and all securities positions held in a segregated account or pursuant to a cross margining agreement, to a regulator (e.g. to the CFTC to meet Part 39, Section 39.19 reporting obligations).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[AccountSummaryReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CR</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitsUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyRiskLimitsUpdateReport(35=CR) is used to convey incremental changes to risk limits. It is similar to the regular report but uses the PartyRiskLimitsUpdateGrp component instead of the PartyRiskLimitsGrp component to include an update action.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyRiskLimitsUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CS</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitsDefinitionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PartyRiskLimitDefinitionRequest is used for defining new risk limits.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [PartyRiskLimitsDefinitionRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CT</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitsDefinitionRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PartyRiskLimitDefinitionRequestAck is used for accepting (with or without changes) or rejecting the definition of risk limits.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">
                                [PartyRiskLimitsDefinitionRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CU</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyEntitlementsRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyEntitlementsRequest message is used to request for entitlement information for one or more party(-ies), specific party role(s), or specific instruments(s).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyEntitlementsRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CV</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyEntitlementsReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyEntitlementsReport is used to report entitlements for one or more parties, party role(s), or specific instrument(s).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyEntitlementsReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CW</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">QuoteAck</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The QuoteAck(35=CW) message is used to acknowledge a Quote(35=S) submittal or request to cancel an individual quote using the QuoteCancel(35=Z) message during a Quote/Negotiation dialog.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[QuoteAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CX</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyDetailsDefinitionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyDetailsDefinitionRequest(35=CX) is used for defining new parties and modifying or deleting existing parties information, including the relationships between parties.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">The recipient of the message responds with a PartyDetailsDefinitionRequestAck(35=CY) to indicate whether the request was accepted or rejected.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyDetailsDefinitionRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">CY</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyDetailsDefinitionRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyDetailsDefinitionRequestAck(35=CY) is used as a response to the PartyDetailsDefinitionRequest(35=CX) message. The request can be accepted (with or without changes) or rejected.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [PartyDetailsDefinitionRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">CZ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyEntitlementsUpdateReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyEntitlementsUpdateReport(35=CZ) is used to convey incremental changes to party entitlements. It is similar to the PartyEntitlementsReport(35=CV). This message uses the PartyEntitlementsUpdateGrp component which includes the ability to specify an update action using ListUpdateAction(1324).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyEntitlementsUpdateReport]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DA</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyEntitlementsDefinitionRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyEntitlementsDefinitionRequest(35=DA) is used for defining new entitlements, and modifying or deleting existing entitlements for the specified party(-ies).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [PartyEntitlementsDefinitionRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DB</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyEntitlementsDefinitionRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyEntitlementsDefinitionRequestAck(35=DB) is used as a response to the PartyEntitlemensDefinitionRequest(35=DA) to accept (with or without changes) or reject the definition of party entitlements.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">
                                [PartyEntitlementsDefinitionRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DC</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradeMatchReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The TradeMatchReport(35=DC) message is used by exchanges and ECN’s to report matched trades to central counterparties (CCPs) as an atomic event. The message is used to express the one-to-one, one-to-many and many-to-many matches as well as implied matches in which more complex instruments can match with simpler instruments.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradeMatchReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DD</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">TradeMatchReportAck</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The TradeMatchReportAck(35=DD) is used to respond to theTradeMatchReport(35=DC) message. It may be used to report on the status of the request (e.g. accepting the request or rejecting the request).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradeMatchReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DE</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitsReportAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PartyRiskLimitsReportAck is an optional message used as a response to the PartyRiskLimitReport(35=CM) or PartyRiskLimitUpdateReport(35=CR) messages to acknowledge or reject those messages.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyRiskLimitsReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DF</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitCheckRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PartyRiskLimitCheckRequest is used to request for approval of credit or risk limit amount intended to be used by a party in a transaction from another party that holds the information.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyRiskLimitCheckRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DG</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PartyRiskLimitCheckRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PartyRiskLimitCheckRequestAck is used to acknowledge a PartyRiskLimitCheckRequest(35=DF) message and to respond whether the limit check request was approved or not. When used to accept the PartyRiskLimitCheckRequest(35=DF) message the Respondent may also include the limit amount that was approved.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyRiskLimitCheckRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DH</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">PartyActionRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PartyActionRequest message is used suspend or halt the specified party from further trading activities at the Respondent. The Respondent must respond with a PartyActionReport(35=DI) message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PartyActionRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DI</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">PartyActionReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">Used to respond to the PartyActionRequest(35=DH) message, indicating whether the request has been received, accepted or rejected. Can also be used in an unsolicited manner to report party actions, e.g. reinstatements after a manual intervention out of band.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PartyActionReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DJ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MassOrder</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The MassOrder(35=DJ) message can be used to add, modify or delete multiple unrelated orders with a single message. Apart from clearing related attributes, only the key order attributes for high performance trading are available.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MassOrder]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DK</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MassOrderAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The mass order acknowledgement message is used to acknowledge the receipt of and the status for a MassOrder(35=DJ) message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MassOrderAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DL</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PositionTransferInstruction</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PositionTransferInstruction(35=DL) is sent by clearing firms to CCPs to initiate position transfers, or to accept or decline position transfers.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PositionTransferInstruction]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DM</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">PositionTransferInstructionAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PositionTransferInstructionAck(35=DM) is sent by CCPs to clearing firms to acknowledge position transfer instructions, and to report errors processing position transfer instructions.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [PositionTransferInstructionAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DN</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PositionTransferReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The PositionTransferReport(35=DN) is sent by CCPs to clearing firms indicating of positions that are to be transferred to the clearing firm, or to report on status of the transfer to the clearing firms involved in the transfer process.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PositionTransferReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DO</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">MarketDataStatisticsRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The MarketDataStatisticsRequest(35=DO) is used to request for statistical data. The simple form is to use an identifier (MDStatisticID(2475)) assigned by the market place which would denote a pre-defined statistical report. Alternatively, or also in addition, the request can define a number of parameters for the desired statistical information.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[MarketDataStatisticsRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DP</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">MarketDataStatisticsReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The MarketDataStatisticsReport(35=DP) is used to provide unsolicited statistical information or in response to a specific request. Each report contains a set of statistics for a single entity which could be a market, a market segment, a security list or an instrument.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarketDataStatisticsReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DQ</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">CollateralReportAck</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">CollateralReportAck(35=DQ) is used as a response to the CollateralReport(35=BA). It can be used to reject a CollateralReport(35=BA) when the content of the report is invalid based on the business rules of the receiver. The message may also be used to acknowledge receipt of a valid CollateralReport(35=BA).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[CollateralReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DR</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">MarketDataReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The MarketDataReport(35=DR) message is used to provide delimiting references (e.g. start and end markers in a continuous broadcast) and details about the number of market data messages sent in a given distribution cycle.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[MarketDataReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DS</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">CrossRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The CrossRequest(35=DS) message is used to indicate the submission of orders or quotes that may result in a crossed trade.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[CrossRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DT</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span
                                        class="">CrossRequestAck</span></p><span
                                        class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">The CrossRequestAck(35=DT) message is used to confirm the receipt of a CrossRequest(35=DS) message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[CrossRequestAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DU</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">AllocationInstructionAlertRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used in a clearinghouse 3-party allocation model to request for AllocationInstructionAlert(35=BM) from the clearinghouse. The request may be used to obtain a one-time notification of the status of an allocation group.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">
                                [AllocationInstructionAlertRequest]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DV</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">AllocationInstructionAlertRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">This message is used in a clearinghouse 3-party allocation model to acknowledge a AllocationInstructionAlertRequest(35=DU) message for an AllocationInstructionAlert(35=BM) message from the clearinghouse.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">
                                [AllocationInstructionAlertRequestAck]
                            </td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DW</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span class="">TradeAggregationRequest</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">TradeAggregationRequest(35=DW) is used to request that the identified trades between the initiator and respondent be aggregated together for further processing.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[TradeAggregationRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DX</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">TradeAggregationReport</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">TradeAggregationReport(35=DX) is used to respond to the TradeAggregationRequest(35=DW) message. It provides the status of the request (e.g. accepted or rejected) and may also provide additional information supplied by the respondent.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[TradeAggregationReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">EA</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">PayManagementReport</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementReport(35=EA) may be used to respond to the PayManagementRequest(35=DY) message. It provides the status of the request (e.g. accepted, disputed) and may provide additional information related to the request.</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementReport(35=EA) may also be sent unsolicited by the broker to a client. In which case the client may acknowledge and resolve disputes out-of-band or with a simple PayManagementReportAck(35=EB).</span>
                                    </p><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementReport(35=EA) may also be sent unsolicited to report the progress status of the payment itself with PayReportTransType(2804)=2 (Status).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PayManagementReport]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">EB</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PayManagementReportAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementReportAck(35=EB) is used as a response to the PayManagementReport(35=EA) message. It may be used to accept, reject or dispute the details of the PayManagementReport(35=EA) depending on the business rules of the receiver. This message may also be used to acknowledge the receipt of a PayManagementReport(35=EA) message.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PayManagementReportAck]</td>
                        </tr>
                        <tr>
                            <td class="EnmTwo">DY</td>
                            <td class="EnmTwo">=</td>
                            <td class="EnmTwo"><span class="">
         </span>
                                <p class=""><span
                                        class="">PayManagementRequest</span></p>
                                <span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementRequest(35=DY) message is used to communicate a future or expected payment to be made or received related to a trade or contract after its settlement.</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo"></td>
                            <td class="EnmTwo">[PayManagementRequest]</td>
                        </tr>
                        <tr>
                            <td class="EnmOne">DZ</td>
                            <td class="EnmOne">=</td>
                            <td class="EnmOne"><span class="">
         </span>
                                <p class=""><span class="">PayManagementRequestAck</span>
                                </p><span class="">
      </span><i><span class="">
         </span>
                                    <p class="Even"><span class="">PayManagementRequestAck(35=DZ) is used to acknowledge the receipt of the PayManagementRequest(35=DY) message (i.e. a technical acknowledgement of receipt). Acceptance or rejection of the request is reported in the corresponding PayManagementReport(35=EA).</span>
                                    </p><span class="">
      </span></i></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne"></td>
                            <td class="EnmOne">[PayManagementRequestAck]</td>
                        </tr>
                    </table>
                </td>
                <td class="FldDeprEven">
                    Added&nbsp;
                    FIX.2.7
                </td>
                <td class="FldDeprEven"></td>
                <td class="FldEnumEven"><a href="tag.html"
                                           target="tagFrame"></a></td>
            </tr>
        </table>
        <hr></hr>
        Used in messages:
        <div class="TagUsedIn"></div>
        <hr></hr>
        Used in components:
        <div class="TagUsedIn">[<A href="body_49485052.html?find=MsgType"
                                   target="mainFrame">StandardHeader</A>]
        </div>
        <br></br>
        <p xmlns=""></p>
        <p xmlns=""></p>
        <table xmlns="">
            <tr>
                <td class="Horiz" id="footer" colspan="2">
                    <ul class="Menu">
                        <li id="copyright">© 2007 - 2017 FIX Protocol Limited
                        </li>
                        <li><a href="http://www.fixtrading.org/contact-us"
                               target="_blank">Contact us</a></li>
                        <li>
                            <a href="http://www.fixtrading.org/terms-and-conditions"
                               target="_blank">Copyright and Acceptable Use
                                policy</a></li>
                        <li><a href="http://www.fixtrading.org/privacy-policy"
                               target="_blank">Privacy Policy</a></li>
                    </ul>
                </td>
            </tr>
        </table>
    </body>
</html>
`