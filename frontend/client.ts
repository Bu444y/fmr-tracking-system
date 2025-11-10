/**
 * FMR Tracking System API Client
 * Clean TypeScript client for the FIPS-compliant Go backend
 * No Encore dependencies
 */

// ============================================================================
// Type Definitions
// ============================================================================

// FMR Types
export type FMRStatus = "Draft" | "Submitted" | "Unresolved" | "In-Progress" | "Resolved" | "Archived";
export type PointOfFailure = "Receipt" | "Testing" | "Mission" | "Other";
export type RepairAction = "RMA" | "Rework" | "Scrap/New Part" | "Other";
export type Signed1149Status = "N/A" | "Waiting for Signature" | "Sign Date";
export type TierLevel = "tier1" | "tier2" | "tier3" | "none";

export interface FMRRequest {
  id: number;
  title: string;
  description?: string;
  status: FMRStatus;
  creationDate: Date;
  resolutionDate?: Date;
  createdAt: Date;
  updatedAt: Date;

  // Report section
  controlNumber?: string;
  program?: string;
  unit?: string;
  ommodSn?: string;
  ommodSystemMeterHrs?: string;
  technicianName?: string;
  email?: string;
  commNumber?: string;
  dateReported?: Date;
  failedDate?: Date;
  pointOfFailure?: PointOfFailure;
  failedPartNomenclature?: string;
  failedPartNumber?: string;
  qty?: string;
  partSerialNumber?: string;
  describeFailure?: string;
  troubleShootingPerformed?: string;
  repairTechnician?: string;
  repairStartDate?: Date;
  repairActivityDescription?: string;

  // Processing section
  repairAction?: RepairAction;
  repairActionOther?: string;

  // Supply Chain section
  signed1149?: Signed1149Status;
  signed1149Date?: Date;
  failedPartReceivedDate?: Date;
  newPartReceivedDate?: Date;
  newPartIssuedDate?: Date;
  newPartSerialNumber?: string;
  supplyComments?: string;

  // Procurement section
  poNumber?: string;
  warrantyStartDate?: Date;
  newPartOrderedDate?: Date;
  vendorContactAddress?: string;
  procurementComments?: string;

  // Logistics section
  requestDate?: Date;
  requestToVendorDate?: Date;
  rmaReceivedDate?: Date;
  shippedToVendorDate?: Date;
  logisticsComment?: string;

  // Contractor section
  problemAssistanceRequested?: string;
  tier1Date?: Date;
  tier2Date?: Date;
  tier3Date?: Date;
  problemResolved?: boolean;
  contractorComments?: string;

  // Additional tracking fields
  formSubmissionType?: string;
  fmrCreatedDate?: Date;
  warrantyStopDate?: Date;
  completedDate?: Date;
  fmrReceivedDateByFailures?: Date;
  trackingNoType?: string;
  trackingNo?: string;
  rmaNoType?: string;
  rmaNo?: string;
  engineerDeployedType?: string;
  engineerName?: string;
  deploymentDate?: Date;
  recordUpdated?: Date;
  recordEnteredOn?: Date;
  recordEnteredBy?: string;
}

export interface CreateFMRRequest {
  title: string;
  description?: string;
  status?: FMRStatus;
  [key: string]: any; // Allow all optional FMR fields
}

export interface UpdateFMRRequest {
  id: number;
  [key: string]: any; // Allow all optional FMR fields
}

export interface FMRListResponse {
  fmrs: FMRRequest[];
  total: number;
}

// Contact Types
export type Team = "The Filius Team" | "CRC AN/TYQ-23A Program Team";

export interface Contact {
  id: number;
  title?: string;
  sequenceNumber?: number;
  firstName: string;
  lastName: string;
  email?: string;
  mobileNumber?: string;
  workNumber?: string;
  organization?: string;
  notes?: string;
  team: Team;
  createdAt: Date;
  updatedAt: Date;
}

export interface CreateContactRequest {
  title?: string;
  sequenceNumber?: number;
  firstName: string;
  lastName: string;
  email?: string;
  mobileNumber?: string;
  workNumber?: string;
  organization?: string;
  notes?: string;
  team: Team;
}

export interface UpdateContactRequest {
  id: number;
  title?: string;
  sequenceNumber?: number;
  firstName?: string;
  lastName?: string;
  email?: string;
  mobileNumber?: string;
  workNumber?: string;
  organization?: string;
  notes?: string;
  team?: Team;
}

export interface ContactsListResponse {
  contacts: Contact[];
  total: number;
}

// Note Types
export type NoteType = "general" | "technician" | "customer";

export interface Note {
  id: number;
  fmrId: number;
  content: string;
  author?: string;
  noteType: NoteType;
  createdAt: Date;
}

export interface CreateNoteRequest {
  fmrId: number;
  content: string;
  author?: string;
  noteType?: NoteType;
}

export interface NotesListResponse {
  notes: Note[];
}

// Preset Types
export interface PresetFilters {
  search?: string;
  statusFilter?: string;
  programFilter?: string;
  unitFilter?: string;
  pointOfFailureFilter?: string;
  repairActionFilter?: string;
  tierLevelFilter?: string;
}

export interface FilterPreset {
  id: number;
  name: string;
  filters: PresetFilters;
  createdAt: Date;
  updatedAt: Date;
}

export interface CreatePresetRequest {
  name: string;
  filters: PresetFilters;
}

export interface ListPresetsResponse {
  presets: FilterPreset[];
}

export interface ColumnPreset {
  id: number;
  name: string;
  columns: string[];
  isDefault: boolean;
  createdAt: Date;
  updatedAt: Date;
}

export interface CreateColumnPresetRequest {
  name: string;
  columns: string[];
  isDefault?: boolean;
}

export interface UpdateColumnPresetRequest {
  id: number;
  name?: string;
  columns?: string[];
  isDefault?: boolean;
}

export interface ListColumnPresetsResponse {
  presets: ColumnPreset[];
}

// Report Types
export interface ExportResponse {
  data: string;
  filename: string;
  contentType: string;
}

// ============================================================================
// Base Client
// ============================================================================

export type BaseURL = string;
export const Local: BaseURL = "http://localhost:4000";

class BaseClient {
  constructor(
    private baseURL: string,
    private options: ClientOptions = {}
  ) {}

  async callAPI(path: string, init?: RequestInit): Promise<Response> {
    const url = `${this.baseURL}${path}`;
    const mergedInit: RequestInit = {
      ...this.options.requestInit,
      ...init,
      headers: {
        "Content-Type": "application/json",
        ...this.options.requestInit?.headers,
        ...init?.headers,
      },
    };

    const fetcher = this.options.fetcher || fetch;
    const response = await fetcher(url, mergedInit);

    if (!response.ok) {
      throw new Error(`API Error: ${response.status} ${response.statusText}`);
    }

    return response;
  }
}

// ============================================================================
// Service Clients
// ============================================================================

export namespace fmr {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async create(params: CreateFMRRequest): Promise<FMRRequest> {
      const resp = await this.baseClient.callAPI(`/fmr`, {
        method: "POST",
        body: JSON.stringify(params),
      });
      return parseJSON(await resp.text());
    }

    async get(params: { id: number }): Promise<FMRRequest> {
      const resp = await this.baseClient.callAPI(`/fmr/${params.id}`, {
        method: "GET",
      });
      return parseJSON(await resp.text());
    }

    async list(params?: {
      status?: FMRStatus;
      search?: string;
      sortBy?: string;
      sortOrder?: "asc" | "desc";
      limit?: number;
      offset?: number;
      tierLevel?: TierLevel;
    }): Promise<FMRListResponse> {
      const queryParams = new URLSearchParams();
      if (params) {
        Object.entries(params).forEach(([key, value]) => {
          if (value !== undefined) {
            queryParams.append(key, String(value));
          }
        });
      }
      const query = queryParams.toString();
      const resp = await this.baseClient.callAPI(
        `/fmr${query ? `?${query}` : ""}`,
        { method: "GET" }
      );
      return parseJSON(await resp.text());
    }

    async update(params: UpdateFMRRequest): Promise<FMRRequest> {
      const { id, ...body } = params;
      const resp = await this.baseClient.callAPI(`/fmr/${id}`, {
        method: "PUT",
        body: JSON.stringify(body),
      });
      return parseJSON(await resp.text());
    }
  }
}

export namespace contacts {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async create(params: CreateContactRequest): Promise<Contact> {
      const resp = await this.baseClient.callAPI(`/contacts`, {
        method: "POST",
        body: JSON.stringify(params),
      });
      return parseJSON(await resp.text());
    }

    async list(): Promise<ContactsListResponse> {
      const resp = await this.baseClient.callAPI(`/contacts`, {
        method: "GET",
      });
      return parseJSON(await resp.text());
    }

    async update(params: UpdateContactRequest): Promise<Contact> {
      const { id, ...body } = params;
      const resp = await this.baseClient.callAPI(`/contacts/${id}`, {
        method: "PUT",
        body: JSON.stringify(body),
      });
      return parseJSON(await resp.text());
    }

    async deleteContact(params: { id: number }): Promise<void> {
      await this.baseClient.callAPI(`/contacts/${params.id}`, {
        method: "DELETE",
      });
    }
  }
}

export namespace notes {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async create(params: CreateNoteRequest): Promise<Note> {
      const resp = await this.baseClient.callAPI(`/notes`, {
        method: "POST",
        body: JSON.stringify(params),
      });
      return parseJSON(await resp.text());
    }

    async list(params: { fmrId: number }): Promise<NotesListResponse> {
      const resp = await this.baseClient.callAPI(`/notes/fmr/${params.fmrId}`, {
        method: "GET",
      });
      return parseJSON(await resp.text());
    }
  }
}

export namespace presets {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async create(params: CreatePresetRequest): Promise<FilterPreset> {
      const resp = await this.baseClient.callAPI(`/presets`, {
        method: "POST",
        body: JSON.stringify(params),
      });
      return parseJSON(await resp.text());
    }

    async list(): Promise<ListPresetsResponse> {
      const resp = await this.baseClient.callAPI(`/presets`, {
        method: "GET",
      });
      return parseJSON(await resp.text());
    }

    async deletePreset(params: { id: number }): Promise<void> {
      await this.baseClient.callAPI(`/presets/${params.id}`, {
        method: "DELETE",
      });
    }
  }
}

export namespace column_presets {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async create(params: CreateColumnPresetRequest): Promise<ColumnPreset> {
      const resp = await this.baseClient.callAPI(`/column-presets`, {
        method: "POST",
        body: JSON.stringify(params),
      });
      return parseJSON(await resp.text());
    }

    async list(): Promise<ListColumnPresetsResponse> {
      const resp = await this.baseClient.callAPI(`/column-presets`, {
        method: "GET",
      });
      return parseJSON(await resp.text());
    }

    async update(params: UpdateColumnPresetRequest): Promise<ColumnPreset> {
      const { id, ...body } = params;
      const resp = await this.baseClient.callAPI(`/column-presets/${id}`, {
        method: "PUT",
        body: JSON.stringify(body),
      });
      return parseJSON(await resp.text());
    }

    async deletePreset(params: { id: number }): Promise<void> {
      await this.baseClient.callAPI(`/column-presets/${params.id}`, {
        method: "DELETE",
      });
    }
  }
}

export namespace reports {
  export class ServiceClient {
    constructor(private baseClient: BaseClient) {}

    async exportData(params?: {
      startDate?: string;
      endDate?: string;
    }): Promise<ExportResponse> {
      const queryParams = new URLSearchParams();
      if (params) {
        Object.entries(params).forEach(([key, value]) => {
          if (value !== undefined) {
            queryParams.append(key, String(value));
          }
        });
      }
      const query = queryParams.toString();
      const resp = await this.baseClient.callAPI(
        `/reports/export${query ? `?${query}` : ""}`,
        { method: "GET" }
      );
      return parseJSON(await resp.text());
    }
  }
}

// ============================================================================
// Client Class
// ============================================================================

export interface ClientOptions {
  fetcher?: (url: string, init?: RequestInit) => Promise<Response>;
  requestInit?: Omit<RequestInit, "headers"> & { headers?: Record<string, string> };
}

export class Client {
  public readonly fmr: fmr.ServiceClient;
  public readonly contacts: contacts.ServiceClient;
  public readonly notes: notes.ServiceClient;
  public readonly presets: presets.ServiceClient;
  public readonly column_presets: column_presets.ServiceClient;
  public readonly reports: reports.ServiceClient;
  private readonly baseClient: BaseClient;
  private readonly target: string;
  private readonly options: ClientOptions;

  constructor(target: BaseURL, options?: ClientOptions) {
    this.target = target;
    this.options = options ?? {};
    this.baseClient = new BaseClient(this.target, this.options);

    this.fmr = new fmr.ServiceClient(this.baseClient);
    this.contacts = new contacts.ServiceClient(this.baseClient);
    this.notes = new notes.ServiceClient(this.baseClient);
    this.presets = new presets.ServiceClient(this.baseClient);
    this.column_presets = new column_presets.ServiceClient(this.baseClient);
    this.reports = new reports.ServiceClient(this.baseClient);
  }

  public with(options: ClientOptions): Client {
    return new Client(this.target, {
      ...this.options,
      ...options,
    });
  }
}

// ============================================================================
// Utilities
// ============================================================================

function parseJSON(text: string): any {
  return JSON.parse(text, dateReviver);
}

function dateReviver(key: string, value: any): any {
  const dateFields = [
    'creationDate', 'resolutionDate', 'createdAt', 'updatedAt',
    'dateReported', 'failedDate', 'repairStartDate', 'signed1149Date',
    'failedPartReceivedDate', 'newPartReceivedDate', 'newPartIssuedDate',
    'warrantyStartDate', 'newPartOrderedDate', 'requestDate',
    'requestToVendorDate', 'rmaReceivedDate', 'shippedToVendorDate',
    'tier1Date', 'tier2Date', 'tier3Date', 'fmrCreatedDate',
    'warrantyStopDate', 'completedDate', 'fmrReceivedDateByFailures',
    'deploymentDate', 'recordUpdated', 'recordEnteredOn'
  ];

  if (dateFields.includes(key) && typeof value === 'string') {
    return new Date(value);
  }
  return value;
}

// ============================================================================
// Default Export
// ============================================================================

const apiURL = import.meta.env.VITE_CLIENT_TARGET || Local;
export default new Client(apiURL, {
  requestInit: { credentials: "include" }
});
